package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/yourfavoritedev/background-changer/helpers"
	"github.com/yourfavoritedev/background-changer/ui"
	"github.com/yourfavoritedev/background-changer/wallpaper"
)

func init() {
	os.Setenv("wallpapersDir", "C:/Users/13236/Documents/wallpapers/")
	// create cache file if it doesn't exist
	_, err := helpers.ReadFile(wallpaper.CurrentWallPaperPath)
	if err != nil {
		err = helpers.WriteFile(wallpaper.CurrentWallPaperPath, []byte{})
		if err != nil {
			return
		}
	}
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
