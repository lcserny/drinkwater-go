package notify

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/go-toast/toast"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	title   = "Drink Water Notification"
	message = "An hour has passed, you need to drink some water!"
)

func OnReady() {
	triggerNotification()

	systray.SetIcon(icon.Data)
	systray.SetTitle("Drink Water!")
	systray.SetTooltip("Drink more water notification app")

	n := newNotifier(1*time.Hour, triggerNotification)

	pauseItem := systray.AddMenuItemCheckbox("Pause", "Pause execution", false)
	go handlePause(pauseItem, n)

	exitItem := systray.AddMenuItem("Exit", "Close the system tray app")
	go handleExit(exitItem)

	n.start()
}

func OnExit() {
	log.Info("Exiting")
}

func handlePause(item *systray.MenuItem, n *notifier) {
	for {
		<-item.ClickedCh
		if item.Checked() {
			item.Uncheck()
			n.unpause()
		} else {
			item.Check()
			n.pause()
		}
	}
}

func triggerNotification() {
	log.Info("Triggered notification")
	notification := toast.Notification{Title: title, Message: message}
	if err := notification.Push(); err != nil {
		log.Error(err)
	}
}

func handleExit(item *systray.MenuItem) {
	<-item.ClickedCh
	systray.Quit()
}
