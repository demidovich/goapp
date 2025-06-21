package files

import (
	"errors"
	"os"
)

func NotExists(path string) bool {
	return !Exists(path)
}

func Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func MakeDirectory(path string) error {
	if Exists(path) {
		return nil
	}

	return os.MkdirAll(path, os.ModePerm)
}

// func MakeFile(path string) error {
// 	if Exists(path) {
// 		return nil
// 	}

// 	file, err := os.Create(path)
// 	file.Close()

// 	return err
// }

func OpenFile(filepath string) (*os.File, error) {
	var file *os.File
	var err error

	if NotExists(filepath) {
		file, err = os.Create(filepath)
	} else {
		file, err = os.Open(filepath)
	}

	return file, err
}

// func WriteFile(filepath string, data []byte) error {
// 	file, err := OpenFile(filepath)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()
// 	_, err = file.Write(data)

// 	return err
// }
