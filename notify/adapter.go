package notify

import (
	"github.com/lxn/walk"
	log "github.com/sirupsen/logrus"
)

const (
	title   = "Drink Water Notification"
	message = "An hour has passed, you need to drink some water!"
)

func CreateMain(title string) *walk.MainWindow {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}
	if err := mw.SetTitle(title); err != nil {
		log.Error(err)
	}
	return mw
}

func CreateTray(mainWindow *walk.MainWindow) *walk.NotifyIcon {
	ni, err := walk.NewNotifyIcon(mainWindow)
	if err != nil {
		log.Fatal(err)
	}
	return ni
}

func AddIconToTray(tray *walk.NotifyIcon, iconPath string) error {
	iconRes, err := walk.Resources.Icon(iconPath)
	if err != nil {
		return err
	}
	return tray.SetIcon(iconRes)
}

func AddPauseContextAction(tray *walk.NotifyIcon, ntf *notifier) error {
	pauseAction := walk.NewAction()
	if err := pauseAction.SetText("Pause"); err != nil {
		return err
	}

	pauseAction.Triggered().Attach(func() {
		if pauseAction.Checked() {
			pauseAction.SetChecked(false)
			ntf.Resume()
		} else {
			pauseAction.SetChecked(true)
			ntf.Pause()
		}
	})

	return tray.ContextMenu().Actions().Add(pauseAction)
}

func AddExitContextAction(tray *walk.NotifyIcon) error {
	exitAction := walk.NewAction()
	if err := exitAction.SetText("Exit"); err != nil {
		return err
	}

	exitAction.Triggered().Attach(func() {
		log.Info("Exiting")
		walk.App().Exit(0)
	})

	return tray.ContextMenu().Actions().Add(exitAction)
}

func TriggerNotification(tray *walk.NotifyIcon) {
	log.Info("Triggered notification")
	if err := tray.ShowInfo(title, message); err != nil {
		log.Error(err)
	}
}
