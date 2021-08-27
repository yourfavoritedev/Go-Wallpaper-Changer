package ps

import (
	"fmt"
	"os"

	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
	"github.com/yourfavoritedev/background-changer/helpers"
)

const (
	PSCmdlet     = "change-background"
	PSScriptPath = "C:/Users/13236/go/src/github.com/yourfavoritedev/background-changer/ps/background-changer.ps1"
	ReplaceLeft  = "$imgPath = "
	ReplaceRight = "\n$code"
)

func ApplyChanges(filename string) {
	b, err := helpers.ReadFile(PSScriptPath)
	if err != nil {
		return
	}

	content := string(b)
	// replace filename path
	currentWallPaperPath := helpers.GetStringInBetween(content, ReplaceLeft, ReplaceRight)
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
