package util

import (
	"io/ioutil"
	"log"
	"os"
)

// LoadFile 获取文件夹当中所有文件名字
func LoadFile(id string,url string) string {
	files, err := ioutil.ReadDir(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name()==id {
			return f.Name()
		}
	}
	return "err"
}

func DelFile(file string) bool {
	image:="./data/file/image/carousel/"+file                 //源文件路径
	err := os.Remove(image)
	if err != nil {
		return false
	} else {
		return true
	}
}
