package worker

import (
	"time"

	"github.com/whym9/receiving_service/pkg/metrics"
	"github.com/whym9/receiving_service/pkg/receiver"
	"github.com/whym9/saving_service/pkg/saver"
)

type Worker struct {
	m  metrics.Metrics
	r  receiver.Receiver
	s1 saver.Saver
	s2 saver.Saver
}

func NewWorker(m metrics.Metrics, r receiver.Receiver, s1, s2 saver.Saver) Worker {
	return Worker{m, r, s1, s2}
}

func (w Worker) Work(metric_addr, addr, dir, dsn string, ch *chan []byte) {
	go w.m.StartMetrics(metric_addr)
	go w.r.StartServer(addr)
	w.s1.Create(dir)
	w.s2.Create(dsn)
	name := time.Now().Format("02-01-2022-59989898")
	for {
		stats := <-*ch

		*ch <- []byte("")

		data := <-*ch

		err := w.s1.Save(data, name)

		if err != nil {
			*ch <- []byte("could not save: " + err.Error())
		}

		err = w.s2.Save(stats, name)

		if err != nil {
			*ch <- []byte("could not save: " + err.Error())
		}
	}

}
