package pinv

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func systrayRun() {
	systray.Run(systrayOnReady, systrayOnExit)
}

func systrayOnReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle(Options.AppName())
	systray.SetTooltip(Options.AppName() + "超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		os.Exit(1)
	}()

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)
}

func systrayOnExit() {
	// clean up here
}
