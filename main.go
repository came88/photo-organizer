package main

import (
	"os"
	"photoorganizer/dropbox"
	"photoorganizer/finder"
	"photoorganizer/group"
	"photoorganizer/model"
)

func main() {

	path, err := dropbox.GetDropboxFolder()
	if err != nil {
		model.Logger.Println("Error: ", err)
		os.Exit(1)
	}
	model.Logger.Println(path)

	path = path + string(os.PathSeparator) + "Camera Uploads"

	photoList, err := finder.FindFiles(path)
	if err != nil {
		model.Logger.Println("Error: ", err)
		os.Exit(1)
	}

	allFiles := group.Group(photoList)

	for key, value := range allFiles {
		model.Logger.Println(key)
		for _, e := range value {
			model.Logger.Println(" ", e)
		}
	}
}
