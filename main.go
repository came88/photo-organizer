package main

import (
	"os"

	. "github.com/came88/photo-organizer/dropbox"
	. "github.com/came88/photo-organizer/finder"
	. "github.com/came88/photo-organizer/group"
	. "github.com/came88/photo-organizer/model"
)

func main() {

	path, err := GetDropboxFolder()
	if err != nil {
		Logger.Println("Error: ", err)
		os.Exit(1)
	}
	Logger.Println(path)

	path = path + string(os.PathSeparator) + "Camera Uploads"

	photoList, err := FindFiles(path)
	if err != nil {
		Logger.Println("Error: ", err)
		os.Exit(1)
	}

	allFiles := By24Hour(photoList)

	count := 0
	for _, value := range allFiles {
		count += len(value)
	}
	Logger.Println("len(photoList) = ", len(photoList))
	Logger.Println("len(value allFiles) = ", count)

	for key, value := range allFiles {
		Logger.Println(key)
		for _, e := range value {
			Logger.Println(" ", e)
		}
		if len(value) < 3 {
			Logger.Println("Non faccio nulla con questo gruppo, ha troppi pochi elementi")
			delete(allFiles, key)
		}
	}

	Logger.Println("---Risultato finale")
	for key, value := range allFiles {
		Logger.Println(key)
		for _, e := range value {
			Logger.Println(" ", e)
		}
	}

	const folderFormat = "2006_01_02"
	Logger.Println("---Inizio a spostare i file")
	for key, value := range allFiles {
		folderName := path + string(os.PathSeparator) + key.Format(folderFormat)
		Logger.Println("Creo la cartella se non esiste", folderName)
		err := os.MkdirAll(folderName, 0o750)
		if err != nil {
			Logger.Println(err)
			Logger.Println("Ignoro questo gruppo")
			continue
		}
		for _, image := range value {
			oldPath := image.BasePath + string(os.PathSeparator) + image.FileName
			newPath := folderName + string(os.PathSeparator) + image.FileName
			Logger.Println("Sposto il file ", oldPath, newPath)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				Logger.Println(err)
			}
		}
	}
}
