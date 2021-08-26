package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/yourfavoritedev/background-changer/ui"
)

func init() {
	os.Setenv("wallpapersDir", "C:/Users/13236/Documents/wallpapers/")
}

func main() {
	// get files from directory
	dir := os.Getenv("wallpapersDir")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read from directory: %v", err)
		return
	}

	// get all file names
	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}

	// initialize termui
	ui.CreateListWidget(fileNames)
}
