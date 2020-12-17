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
// fino al massimo alle ore 6 del giorno dopo se non ci sono periodi senza file
// di più di 3 ore
func By24Hour(photoList []Image) map[time.Time][]Image {
	allFiles := make(map[time.Time][]Image)

	// photoList deve essere ordinato
	sort.SliceStable(photoList, func(i, j int) bool {
		return photoList[i].FileName < photoList[j].FileName
	})

	day, _ := time.ParseDuration("24h")
	maxEmpty, _ := time.ParseDuration("3h")

	for {
		Logger.Println("len(photoList) = ", len(photoList))
		if len(photoList) == 0 {
			break
		}
		Logger.Println(photoList[0])
		timeStart := photoList[0].Time
		currentDay := midnight(timeStart)
		currentList := make([]Image, 0)
		currentList = append(currentList, photoList[0])
		for i, currentImage := range photoList[1:] {
			Logger.Println(i+1, currentImage)
			difference := currentImage.Time.Sub(timeStart)
			Logger.Println("difference: ", difference)
			if difference <= day {
				if midnight(currentImage.Time) != currentDay {
					if currentImage.Time.Hour() >= 6 {
						// non aggiungo se l'orario è successivo alle 6 di mattino
						break
					}
					lastDifference := currentImage.Time.Sub(currentList[len(currentList)-1].Time)
					if lastDifference > maxEmpty {
						break
					}

				}
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
