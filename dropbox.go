package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
)

// GetDropboxFolder restituisce la cartella dropbox
func GetDropboxFolder() (string, error) {

	if runtime.GOOS == "windows" {
		// cerco il file info.json in %LOCALAPPDATA%/Dropbox
		localappdata := os.Getenv("LOCALAPPDATA")
		if localappdata != "" {
			nomefile := localappdata + "/Dropbox/info.json"
			if fileExists(nomefile) {
				folder, err := readJSON(nomefile)
				if err == nil {
					return folder, nil
				}
			}
		}
		// cerco il file info.json in %APPDATA%/Dropbox
		appdata := os.Getenv("APPDATA")
		if appdata != "" {
			nomefile := appdata + "/Dropbox/info.json"
			if fileExists(nomefile) {
				folder, err := readJSON(nomefile)
				if err == nil {
					return folder, nil
				}
			}
		}
	}

	// non-Windows
	// cerco il file info.json in ~/.dropbox
	usr, err := user.Current()
	if err != nil {
		return "", errors.New("unable to find Dropbox folder")
	}
	nomefile := usr.HomeDir + ".dropbox/info.json"
	if fileExists(nomefile) {
		folder, err := readJSON(nomefile)
		if err == nil {
			return folder, nil
		}
	}
	return "", errors.New("unable to find Dropbox folder")
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readJSON(nomefile string) (string, error) {

	jsonFile, err := os.Open(nomefile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return "", err
	}

	personal, ok := result["personal"].(map[string]interface{})

	if ok != true {
		return "", errors.New("invalid json format")
	}

	var path string

	path, ok = personal["path"].(string)

	if ok != true {
		return "", errors.New("invalid json format")
	}

	return path, nil
}
