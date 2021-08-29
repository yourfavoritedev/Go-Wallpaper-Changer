package wallpaper

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/yourfavoritedev/background-changer/helpers"
)

const (
	CurrentWallPaperPath = "C:/Users/13236/go/src/github.com/yourfavoritedev/background-changer/cache/current-wallpaper.txt"
)

// UI Params for windows
const (
	spiGetdeskwallpaper = 0x0073
	spiSetdeskwallpaper = 0x0014

	uiParam = 0x0000

	pifUpdateINIFile = 0x01
	spifSendChange   = 0x02
)

// user32.dll and its proc
var (
	user32                = syscall.NewLazyDLL("user32.dll")
	systemParametersInfoW = user32.NewProc("SystemParametersInfoW")
)

func ApplyChanges(filename string) {
	// construct filename path
	newWallPaperPath := fmt.Sprintf("%s%s", os.Getenv("wallpapersDir"), filename)
	helpers.WriteFile(CurrentWallPaperPath, []byte(newWallPaperPath))
	SetWallpaper(newWallPaperPath)
}

func SetWallpaper(filename string) error {
	filenameUTF16Ptr, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}

	systemParametersInfoW.Call(
		uintptr(spiSetdeskwallpaper),              // Message
		uintptr(uiParam),                          // UI Param
		uintptr(unsafe.Pointer(filenameUTF16Ptr)), // User argument e.g. file name
		uintptr(pifUpdateINIFile|spifSendChange),  // we want to update the user profile and set this change into registry
	)

	return nil
}
