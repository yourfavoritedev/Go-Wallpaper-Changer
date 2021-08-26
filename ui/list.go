package ui

import (
	"log"

	term "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/yourfavoritedev/background-changer/ps"
)

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
	l.SetRect(0, 0, 25, 8)

	term.Render(l)
	uiEvents := term.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "w":
			l.ScrollUp()
		case "s":
			l.ScrollDown()
		}
		term.Render(l)

		// update power-shell script file
		selectedIndex := l.SelectedRow
		selectedImageFilename := list[selectedIndex]
		ps.ApplyChanges(selectedImageFilename)
	}
}
