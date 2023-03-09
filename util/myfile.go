package util

import (
	"io/ioutil"
	"log"
	"os"
)

// LoadFile 获取文件夹当中所有文件名字
func LoadFile(id string, url string) string {
	files, err := ioutil.ReadDir(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() == id {
			return f.Name()
		}
	}
	return "err"
}

// LoadFileList 获取文件列表名字
func LoadFileList(url string) []string {
	files, err := ioutil.ReadDir(url)
	if err != nil {
		log.Fatal(err)
	}
	var FileList []string = make([]string, 0)
	for _, f := range files {
		FileList = append(FileList, f.Name())
	}
	return FileList
}

func DelFile(file string) bool {
	image := "./data/file/image/carousel/" + file //源文件路径
	err := os.Remove(image)
	if err != nil {
		return false
	} else {
		return true
	}
}
