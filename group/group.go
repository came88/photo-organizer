package group

import (
	. "github.com/came88/photo-organizer/model"
	"time"
)

func midnight(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
}

func Group(photoList []Image) map[time.Time][]Image {
	allFiles := make(map[time.Time][]Image)

	for _, image := range photoList {
		currentList := allFiles[midnight(image.Time)]
		currentList = append(currentList, image)
		allFiles[midnight(image.Time)] = currentList
	}

	return allFiles
}
