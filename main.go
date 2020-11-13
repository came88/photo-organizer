package main

import (
	"os"
	. "photoorganizer/dropbox"
	. "photoorganizer/finder"
	. "photoorganizer/group"
	. "photoorganizer/model"
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

	allFiles := Group(photoList)

	for key, value := range allFiles {
		Logger.Println(key)
		for _, e := range value {
			Logger.Println(" ", e)
		}
	}
}
