package model

import "time"

type Image struct {
	BasePath string
	FileName string
	Time     time.Time
}
