package main

import (
	"flag"

	"github.com/whym9/saving_service/pkg/server"
)

func main() {
	dsn := *flag.String("dsn", "root:Rarefictions5@tcp(127.0.0.1:3306)/pcap_files?charset=utf8mb4", "dsn")
	addr := *flag.String("address", "localhost:443", "server address")
	dir := *flag.String("dir", "files", "directory for saving files")
	flag.Parse()

	server.Server(addr, dir, dsn)
}
