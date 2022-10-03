package saver

type Saver interface {
	Save(counter []byte, filepath string) error
}
