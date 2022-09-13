package worker

import (
	"time"

	"github.com/whym9/receiving_service/pkg/metrics"
	"github.com/whym9/receiving_service/pkg/receiver"
	"github.com/whym9/saving_service/pkg/saver"
)

type Worker struct {
	m metrics.Metrics
	r receiver.Receiver
	s saver.Saver
}

func NewWorker(m metrics.Metrics, r receiver.Receiver, s saver.Saver) Worker {
	return Worker{m, r, s}
}

func (w Worker) Work(metric_addr, addr, dir, dsn string, ch chan []byte) {
	go w.m.StartMetrics(metric_addr)
	go w.r.StartServer(addr)
	w.s.Create(dsn)

	name := dir + time.Now().Format("02-01-2022-59989898")
	for {

		data := <-ch

		err := w.s.Save(data, name)

		if err != nil {
			ch <- []byte("could not save: " + err.Error())
		}

	}

}
