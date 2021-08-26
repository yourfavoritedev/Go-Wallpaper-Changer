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
	psCmdlet     = "change-background"
	psScriptPath = "C:/Users/13236/go/src/github.com/yourfavoritedev/background-changer/ps/background-changer.ps1"
	replaceStart = "$imgPath = "
	replaceEnd   = "$code"
)

func ApplyChanges(filename string) {
	b, err := helpers.ReadFile(psScriptPath)
	if err != nil {
		return
	}

	content := string(b)
	// replace filename path
	s := strings.Index(content, replaceStart)
	s += len(replaceStart)
	e := strings.Index(content[s:], replaceEnd)
	e += s - 1
	currentWallPaperPath := content[s:e]
	newWallPaperPath := fmt.Sprintf("%s%s%s%s", "\"", os.Getenv("wallpapersDir"), filename, "\"")
	updatedContent := helpers.ReplaceText(content, currentWallPaperPath, newWallPaperPath)
	err = helpers.WriteFile(psScriptPath, []byte(updatedContent))
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
	_, _, err = shell.Execute(psCmdlet)
	if err != nil {
		panic(err)
	}
}
