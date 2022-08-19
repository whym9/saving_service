package main

import (
	"flag"
	"fmt"
	"net"
	"saving_service/internal/metrics"
	sendpb "saving_service/internal/proto"
	"saving_service/internal/saver"
	"saving_service/internal/server"
	"saving_service/internal/storage"

	"log"

	"google.golang.org/grpc"
)

func main() {
	dsn := *flag.String("dsn", "root:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4", "dsn")
	addr := *flag.String("address", "localhost:443", "server address")
	dir := *flag.String("dir", "files", "directory for saving files")
	flag.Parse()
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	uplSrv := server.NewServer(storage.New(""), dir)

	rpcSrv := grpc.NewServer()

	handle := saver.DB_Handle{}

	err = handle.CreateDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	uplSrv.Handle = handle
	go metrics.Metrics(addr)
	fmt.Println("GRPC server has started")
	sendpb.RegisterSendServiceServer(rpcSrv, uplSrv)

	log.Fatal(rpcSrv.Serve(lis))
}
