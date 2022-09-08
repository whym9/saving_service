package file

import "os"

type file_Handle struct {
	dir string
}

func NewFileHandle() file_Handle {
	return file_Handle{}
}

func (f file_Handle) Create(dir string) error {
	err := os.MkdirAll(dir, os.ModeAppend)

	if err != nil {
		return err
	}
	f.dir = dir
	return nil
}

func (f file_Handle) Save(data []byte, name string) error {
	file, err := os.OpenFile(f.dir+"/"+name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(data)
	return nil
}
