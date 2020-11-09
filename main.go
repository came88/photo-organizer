package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path, err := GetDropboxFolder()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(path)

	path = path + string(os.PathSeparator) + "Camera Uploads"

	files, err := ioutil.ReadDir(path)

	for _, file := range files {
		fmt.Println(file.Name())
	}


}
