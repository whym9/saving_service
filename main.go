package main

import (
	metrics "github.com/whym9/receiving_service/pkg/metrics/prometheus"
	receiver "github.com/whym9/receiving_service/pkg/receiver/GRPC"

	"github.com/whym9/saving_service/internal/worker"

	saver "github.com/whym9/saving_service/pkg/saver/mySql"
)

func main() {

	ch := make(chan []byte)
	Promo_Handler := metrics.NewPromoHandler()

	mysql_saver := saver.NewDBHandle(Promo_Handler)
	GRPC_server := receiver.NewServer(Promo_Handler, ch)

	w := worker.NewWorker(Promo_Handler, GRPC_server, mysql_saver)

	w.Work(ch)
}
