package model

import (
	"log"
	"os"
	"time"
)

// Image rappresenta un immagine con cartella, nome file e data
type Image struct {
	BasePath string
	FileName string
	Time     time.Time
}

// Logger per tutto il programma
var Logger = log.New(os.Stdout, "photoorganizer: ", log.Lshortfile)
