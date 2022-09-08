package saver

type Saver interface {
	Create(name string) error
	Save(counter []byte, filepath string) error
}
