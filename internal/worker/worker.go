package worker

import (
	"fmt"
	"strings"
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

func (w Worker) Work(ch chan []byte) {
	go w.m.StartMetrics()
	go w.r.StartServer()

	for {

		data := <-ch
		fmt.Println(len(data))
		name := "files/" + time.Now().Format("2006.01.02 15:04:05") + ".pcap"
		name = strings.Replace(name, " ", "-", -1)
		name = strings.Replace(name, ":", ".", -1)
		fmt.Println(name)
		err := w.s.Save(data, name)

		if err != nil {
			fmt.Println(err.Error())
			ch <- []byte("could not save: " + err.Error())
		}
		ch <- []byte("")

	}

}
