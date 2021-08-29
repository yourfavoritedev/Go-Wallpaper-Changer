package ui

import (
	"log"
	"os"

	term "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/yourfavoritedev/background-changer/helpers"
	"github.com/yourfavoritedev/background-changer/wallpaper"
)

func GetInitialFilename() (result string) {
	b, err := helpers.ReadFile(wallpaper.CurrentWallPaperPath)
	if err != nil {
		return
	}
	content := string(b)
	wallPaperDir := os.Getenv("wallpapersDir")
	currentFileName := content[len(wallPaperDir):]
	return currentFileName
}

func CreateListWidget(list []string) {
	if err := term.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer term.Close()

	// construct list
	l := widgets.NewList()
	l.Title = "Images"
	l.Rows = list
	l.TextStyle = term.NewStyle(term.ColorYellow)
	l.WrapText = false
	initialFilename := GetInitialFilename()
	for i := range list {
		if list[i] == initialFilename {
			l.SelectedRow = i
			break
		}
	}

	l.SetRect(0, 0, 25, 8)

	term.Render(l)
	uiEvents := term.PollEvents()

	// enables wrap-around list
	previousSelectedRow := l.SelectedRow

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "w":
			if previousSelectedRow == 0 {
				l.SelectedRow = len(list) - 1
			} else {
				l.ScrollUp()
			}

		case "s":
			if previousSelectedRow == len(list)-1 {
				l.SelectedRow = 0
			} else {
				l.ScrollDown()
			}
		}
		previousSelectedRow = l.SelectedRow
		term.Render(l)

		// update power-shell script file
		selectedIndex := l.SelectedRow
		selectedImageFilename := list[selectedIndex]
		wallpaper.ApplyChanges(selectedImageFilename)
	}
}
