package model

import (
	"log"
	"os"
	"time"
)

type Image struct {
	BasePath string
	FileName string
	Time     time.Time
}

var Logger = log.New(os.Stdout, "photoorganizer: ", log.Lshortfile)
