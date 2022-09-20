package saver

type Saver interface {
	Create() error
	Save(counter []byte, filepath string) error
}
