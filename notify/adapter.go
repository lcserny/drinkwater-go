package notify

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/go-toast/toast"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	notificationDelay = 1 * time.Hour

	appId   = "DrinkwaterGo" // when changing this please also update the build script
	title   = "Drink Water Notification"
	message = "An hour has passed, you need to drink some water!"
)

func OnReady() {
	triggerNotification()

	systray.SetIcon(icon.Data)
	systray.SetTitle("Drink Water!")
	systray.SetTooltip("Drink more water notification app")

	n := newNotifier(notificationDelay, triggerNotification)
	pauseItem := systray.AddMenuItemCheckbox("Pause", "Pause execution", false)
	exitItem := systray.AddMenuItem("Exit", "Close the system tray app")
	n.start()

	listenForCommands(n, pauseItem, exitItem)
}

func listenForCommands(n *notifier, pauseItem *systray.MenuItem, exitItem *systray.MenuItem) {
	for {
		select {
		case <-pauseItem.ClickedCh:
			if pauseItem.Checked() {
				pauseItem.Uncheck()
				n.resume()
			} else {
				pauseItem.Check()
				n.pause()
			}
		case <-exitItem.ClickedCh:
			systray.Quit()
		}
	}
}

func OnExit() {
	log.Info("Exiting")
}

func triggerNotification() {
	log.Info("Triggered notification")
	notification := toast.Notification{AppID: appId, Title: title, Message: message}
	if err := notification.Push(); err != nil {
		log.Error(err)
	}
}
