package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Drink Water!")
	systray.SetTooltip("Drink more water notification app")

	pauseItem := systray.AddMenuItemCheckbox("Pause", "Pause execution", false)
	go handlePause(pauseItem)

	// TODO: remove once scheduled job is ok
	notifyItem := systray.AddMenuItem("Notify", "Show popup notification")
	go handleNotify(notifyItem)

	exitItem := systray.AddMenuItem("Exit", "Exit the whole app")
	go handleExit(exitItem)

	// TODO: add scheduled job with auto-trigger notify, use channel for outside pause command
}

func onExit() {
	fmt.Println("Exiting")
}

// TODO: stop scheduled job, use channels
func handlePause(item *systray.MenuItem) {
	for {
		<-item.ClickedCh
		if item.Checked() {
			item.Uncheck()
		} else {
			item.Check()
		}
		fmt.Println("Clicked pause, value to:", item.Checked())
	}
}

func handleNotify(item *systray.MenuItem) {
	for {
		<-item.ClickedCh
		triggerNotification()
		fmt.Println("Clicked notify")
	}
}

// TODO: add image in code and use go embed
func triggerNotification() {
	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		// TODO: implement logging
		panic(err)
	}
}

func handleExit(item *systray.MenuItem) {
	<-item.ClickedCh
	systray.Quit()
}
