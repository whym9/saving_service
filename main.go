package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"saving_service/process"
	sendpb "saving_service/proto"
	"saving_service/saver"
	"saving_service/storage"

	"time"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	dsn := *flag.String("dsn", "root:password5@tcp(127.0.0.1:3306)/db?charset=utf8mb4", "dsn")
	addr := *flag.String("address", "localhost:443", "server address")
	flag.Parse()
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	uplSrv := NewServer()

	rpcSrv := grpc.NewServer()

	handle := saver.DB_Handle{}

	err = handle.CreateDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	uplSrv.handle = handle

	fmt.Println("GRPC server has started")
	sendpb.RegisterSendServiceServer(rpcSrv, uplSrv)

	log.Fatal(rpcSrv.Serve(lis))
}

type Server struct {
	sendpb.UnimplementedSendServiceServer
	handle saver.DB_Handle
}

func NewServer() Server {
	return Server{}
}

func (s Server) Send(stream sendpb.SendService_SendServer) error {

	dirName := "files"
	err := os.MkdirAll(dirName, os.ModePerm)

	if err != nil {

		return err
	}
	counter := process.Protocols{}

	name := dirName + "/" + time.Now().Format("02-01-1975-547896") + ".pcapng"

	f := storage.NewFile(name)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	req, err := stream.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	err = json.Unmarshal(req.GetChunk(), &counter)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	packets := []process.Packet{}
	for {

		req, err := stream.Recv()

		if err == io.EOF {
			err := f.Write(packets)
			if err != nil {

				return status.Error(codes.Internal, err.Error())
			}
			err = s.handle.SaveToDB(counter, name)

			if err != nil {

				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&sendpb.SendResponse{Name: "ok"})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		cap := process.Capture{}
		bin := req.GetChunk()

		err = json.Unmarshal(bin, &cap)

		req, err = stream.Recv()

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		packets = append(packets, process.NewPacket(cap, req.GetChunk()))

	}

}
