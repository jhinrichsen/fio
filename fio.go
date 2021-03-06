package fio

import (
	"log"
	"os"
	"time"
)

// Exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	fi, err := os.Stat(path)
	log.Printf("stat %s: %+v\n", path, fi)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	var mu bool
	return mu, err
}

func ModTime(filename string) (time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return time.Now(), err
	}
	return fi.ModTime(), nil
}
