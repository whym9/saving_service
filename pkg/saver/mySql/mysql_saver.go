package saver

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/whym9/receiving_service/pkg/metrics"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Protocols struct {
	TCP  int `json: "TCP"`
	UDP  int `json: "UDP"`
	IPv4 int `json: "IPv4"`
	IPv6 int `json: "IPv6"`
}

type DB_Handle struct {
	metrics metrics.Metrics
	DB      *gorm.DB
	Name    string
	stats   file_statistics
}

func NewDBHandle(m metrics.Metrics) DB_Handle {
	dsn := os.Getenv("DSN")

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}
	m.AddMetrics(name, help, key)
	return DB_Handle{metrics: m, DB: DB}
}

type file_statistics struct {
	Filepath    string
	Protocoltcp int
	UDP         int
	IPv4        int
	IPv6        int
}

var (
	name = "Errors_in_saving_to_DB_total"
	help = "The total number of errors in saving to DB"
	key  = "errors"
)

func (h DB_Handle) Save(data []byte, filePath string) error {

	if err := json.Unmarshal(data, &Protocols{}); err == nil {

		counter := Protocols{}
		err := json.Unmarshal(data, &counter)
		if err != nil {
			fmt.Println(err.Error())
			h.metrics.Count(key)
			return err
		}
		result := h.DB.Create(&file_statistics{

			Filepath:    filePath,
			Protocoltcp: counter.TCP,
			UDP:         counter.UDP,
			IPv4:        counter.IPv4,
			IPv6:        counter.IPv6,
		})

		if result.Error != nil {
			h.metrics.Count(key)
			fmt.Println(err.Error())
			return result.Error
		}
		fmt.Println()
		fmt.Println("Record saved to Database!")
		fmt.Println()
		return nil
	}

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("File uploaded")

	return nil

}
