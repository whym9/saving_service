package main

import (
	"flag"

	metrics "github.com/whym9/receiving_service/pkg/metrics/prometheus"
	receiver "github.com/whym9/receiving_service/pkg/receiver/GRPC"

	"github.com/whym9/saving_service/internal/worker"
	saver1 "github.com/whym9/saving_service/pkg/saver/file"
	saver2 "github.com/whym9/saving_service/pkg/saver/mySql"
)

func main() {
	dsn := *flag.String("dsn", "root:Rarefictions5@tcp(127.0.0.1:3306)/pcap_files?charset=utf8mb4", "dsn")
	addr := *flag.String("address", "localhost:443", "server address")
	dir := *flag.String("dir", "files", "directory for saving files")
	metric_addr := *flag.String("metric_addr", "7007", "address where to run metrics server")
	flag.Parse()
	ch := make(chan []byte)
	Promo_Handler := metrics.NewPromoHandler()

	file_saver := saver1.NewFileHandle()

	mysql_saver := saver2.NewDBHandle(Promo_Handler)
	GRPC_server := receiver.NewServer(Promo_Handler, &ch)

	w := worker.NewWorker(Promo_Handler, GRPC_server, file_saver, mysql_saver)

	w.Work(metric_addr, addr, dir, dsn, &ch)
}
