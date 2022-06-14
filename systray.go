package pinv

import (
	"os"
	"os/exec"
	"syscall"

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
	mUrl := systray.AddMenuItem("Open", "open url")
	go func() {
		<-mUrl.ClickedCh
		//windows
		cmd := exec.Command("cmd", "/c", "start", "http://127.0.0.1:8080")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Start()
	}()
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
