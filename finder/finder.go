package finder

import (
	"io/ioutil"
	"time"

	. "github.com/came88/photo-organizer/model"
)

func FindFiles(path string) ([]Image, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// deve rappresentare il formato desiderato per la data "Mon Jan 2 15:04:05 -0700 MST 2006"
	const dateLayout = "2006-01-02 15.04.05"

	photoList := make([]Image, 0)

	for _, file := range files {
		if file.IsDir() || len(file.Name()) < len(dateLayout) {
			continue
		}
		date, err := time.Parse(dateLayout, file.Name()[:len(dateLayout)])
		if err == nil {
			image := Image{
				BasePath: path,
				FileName: file.Name(),
				Time:     date,
			}
			photoList = append(photoList, image)
		} else {
			Logger.Println(err)
		}
	}

	return photoList, nil
}
