package pinv

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/skratchdot/open-golang/open"
)

func systrayRun() {
	systray.Run(systrayOnReady, systrayOnExit)
}

func systrayOnReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle(Options.AppName())
	systray.SetTooltip(Options.AppName() + " " + Options.Version())
	mUrl := systray.AddMenuItem("Open", "open url")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)

	go func() {
		for {
			select {
			case <-mUrl.ClickedCh:
				open.Run("http://127.0.0.1:8080")
			case <-mQuit.ClickedCh:
				systray.Quit()
				os.Exit(1)
			}
		}
	}()

}

func systrayOnExit() {
	// clean up here
}
