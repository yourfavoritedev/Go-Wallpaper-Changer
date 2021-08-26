package ps

import (
	"fmt"
	"os"
	"strings"

	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
	"github.com/yourfavoritedev/background-changer/helpers"
)

const (
	PSCmdlet     = "change-background"
	PSScriptPath = "C:/Users/13236/go/src/github.com/yourfavoritedev/background-changer/ps/background-changer.ps1"
	ReplaceStart = "$imgPath = "
	ReplaceEnd   = "$code"
)

func ApplyChanges(filename string) {
	b, err := helpers.ReadFile(PSScriptPath)
	if err != nil {
		return
	}

	content := string(b)
	// replace filename path
	s := strings.Index(content, ReplaceStart)
	s += len(ReplaceStart)
	e := strings.Index(content[s:], ReplaceEnd)
	e += s - 1
	currentWallPaperPath := content[s:e]
	newWallPaperPath := fmt.Sprintf("%s%s%s%s", "\"", os.Getenv("wallpapersDir"), filename, "\"")
	updatedContent := helpers.ReplaceText(content, currentWallPaperPath, newWallPaperPath)
	err = helpers.WriteFile(PSScriptPath, []byte(updatedContent))
	if err != nil {
		return
	}

	RunPowershell()
}

func RunPowershell() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... and interact with it
	_, _, err = shell.Execute(PSCmdlet)
	if err != nil {
		panic(err)
	}
}
