package saver

import (
	"fmt"

	"saving_service/process"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Saver interface {
	CreateDB(dsn string) error
	SaveToDB(counter process.Protocols, filepath string) error
}

type DB_Handle struct {
	DB *gorm.DB
}

type file_statistics struct {
	Filepath    string
	Protocoltcp int
	UDP         int
	IPv4        int
	IPv6        int
}

func (h *DB_Handle) CreateDB(dsn string) error {
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error")
		return err
	}
	h.DB = DB
	return nil
}

func (h DB_Handle) SaveToDB(counter process.Protocols, filePath string) error {

	result := h.DB.Create(&file_statistics{

		Filepath:    filePath,
		Protocoltcp: counter.TCP,
		UDP:         counter.UDP,
		IPv4:        counter.IPv4,
		IPv6:        counter.IPv6,
	})

	if result.Error != nil {

		return result.Error
	}
	fmt.Println()
	fmt.Println("Record saved to Database!")
	fmt.Println()
	return nil
}
