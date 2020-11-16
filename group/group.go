package group

import (
	"sort"
	"time"

	. "github.com/came88/photo-organizer/model"
)

func midnight(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

// BySolarDay raggruppa per giorno solare
func BySolarDay(photoList []Image) map[time.Time][]Image {
	allFiles := make(map[time.Time][]Image)

	for _, image := range photoList {
		currentList := allFiles[midnight(image.Time)]
		currentList = append(currentList, image)
		allFiles[midnight(image.Time)] = currentList
	}

	// ordino le sottoliste sopo averle splittate, è più veloce
	for _, valueList := range allFiles {
		sort.SliceStable(valueList, func(i, j int) bool {
			return valueList[i].FileName < valueList[j].FileName
		})
	}

	return allFiles
}

// By24Hour raggruppa per periodi di 24 ore, partendo dalla prima occorrenza
func By24Hour(photoList []Image) map[time.Time][]Image {
	allFiles := make(map[time.Time][]Image)

	// photoList deve essere ordinato
	sort.SliceStable(photoList, func(i, j int) bool {
		return photoList[i].FileName < photoList[j].FileName
	})

	day, _ := time.ParseDuration("24h")

	for {
		Logger.Println("len(photoList) = ", len(photoList))
		if len(photoList) == 0 {
			break
		}
		Logger.Println(photoList[0])
		timeStart := photoList[0].Time
		currentList := make([]Image, 0)
		currentList = append(currentList, photoList[0])
		for i, currentImage := range photoList[1:] {
			Logger.Println(i+1, currentImage)
			difference := currentImage.Time.Sub(timeStart)
			Logger.Println("difference: ", difference)
			if difference <= day {
				currentList = append(currentList, currentImage)
			} else {
				break
			}
		}
		photoList = photoList[len(currentList):]
		Logger.Println("aggiungo ", currentList)
		allFiles[midnight(timeStart)] = currentList
	}

	return allFiles
}
