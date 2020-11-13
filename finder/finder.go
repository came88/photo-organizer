package finder

import (
	"io/ioutil"
	"photoorganizer/model"
	"sort"
	"time"
)

// TODO la ricerca fa fatta ricorsiva?
func FindFiles(path string) ([]model.Image, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// deve rappresentare la data "02/Jan/2006:15:04:05
	const dateLayout = "2006-01-02 15.04.05"

	photoList := make([]model.Image, 1)

	for _, file := range files {
		if file.IsDir() || len(file.Name()) < len(dateLayout) {
			continue
		}
		date, err := time.Parse(dateLayout, file.Name()[:len(dateLayout)])
		if err == nil {
			image := model.Image{
				BasePath: path,
				FileName: file.Name(),
				Time:     date,
			}
			photoList = append(photoList, image)
		} else {
			model.Logger.Println(err)
		}
	}

	sort.SliceStable(photoList, func(i, j int) bool {
		return photoList[i].FileName < photoList[j].FileName
	})

	return photoList, nil
}
