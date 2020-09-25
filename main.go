package main

import (
	"fmt"
)

func main() {
	path, err := GetDropboxFolder()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}
}
