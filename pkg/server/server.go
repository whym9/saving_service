package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/whym9/receiving_service/pkg/metrics"
	"github.com/whym9/receiving_service/pkg/receiver"
	"github.com/whym9/receiving_service/pkg/saver"
)

func Server(addr, dir, dsn string) {
	fmt.Println("Start")
	go metrics.PromoHandler{}.StartMetrics("6060")
	err := os.MkdirAll(dir, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan []byte)
	s := receiver.NewServer(&ch)

	go s.StartServer(addr)
	for {
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

}
