package server

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/whym9/receiving_service/pkg/metrics"
	"github.com/whym9/receiving_service/pkg/receiver"
	"github.com/whym9/saving_service/pkg/saver"
)

func Server(addr, dir, dsn string) {
	metrics.PromoHandler{}.StartMetrics(addr)
	err := os.MkdirAll(dir, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan []byte)
	s := receiver.NewServer(&ch)

	go s.StartServer(addr)

	bin := <-ch

	counter := saver.Protocols{}

	err = json.Unmarshal(bin, &counter)

	ch <- []byte("")

	name := dir + "/" + time.Now().Format("02-01-2002-5945995")
	file, err := os.Create(name)
	defer file.Close()

	file.Write(<-ch)
	h := saver.DB_Handle{}
	err = h.CreateDB(dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = h.SaveToDB(counter, name)

	if err != nil {
		ch <- []byte("could not save")
	}

	ch <- []byte("")
}
