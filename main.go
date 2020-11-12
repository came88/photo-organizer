package main

import (
	"fmt"
	"os"
	"photoorganizer/dropbox"
	"photoorganizer/finder"
	"photoorganizer/group"
)

func main() {
	path, err := dropbox.GetDropboxFolder()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(path)

	path = path + string(os.PathSeparator) + "Camera Uploads"
	photoList, err := finder.FindFiles(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	allFiles := group.Group(photoList)

	for key, value := range allFiles {
		fmt.Println(key)
		for _, e := range value {
			fmt.Println(" ", e)
		}
	}
}
