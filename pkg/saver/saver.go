package saver

type Saver interface {
	CreateDB(dsn string) error
	SaveToDB(counter Protocols, filepath string) error
}
