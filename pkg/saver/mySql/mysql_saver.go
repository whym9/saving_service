package saver

import (
	"encoding/json"
	"fmt"

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
}

func NewDBHandle(m metrics.Metrics) DB_Handle {
	return DB_Handle{metrics: m}
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

func (h DB_Handle) Create(dsn string) error {
	h.metrics.AddMetrics(name, help, key)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		h.metrics.Count(key)
		return err
	}
	h.DB = DB
	return nil
}

func (h DB_Handle) Save(data []byte, filePath string) error {
	counter := Protocols{}
	err := json.Unmarshal(data, &counter)
	if err != nil {
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
		return result.Error
	}
	fmt.Println()
	fmt.Println("Record saved to Database!")
	fmt.Println()
	return nil
}