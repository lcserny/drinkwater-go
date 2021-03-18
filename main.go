package main

import (
	"drinkwater-go/notify"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	logFile           = "drinkwater.log"
	notificationDelay = 1 * time.Hour
)

func initLogging() *os.File {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}

func main() {
	f := initLogging()
	defer f.Close()

	mainWindow := notify.CreateMain("Drink Water Notification")

	tray := notify.CreateTray(mainWindow)
	defer tray.Dispose()

	if err := notify.AddIconToTray(tray, "glass_original.ico"); err != nil {
		log.Error(err)
	}

	if err := tray.SetToolTip("Drink more water notification app"); err != nil {
		log.Fatal(err)
	}

	ntf := notify.NewNotifier(notificationDelay, func() { notify.TriggerNotification(tray) })

	if err := notify.AddPauseContextAction(tray, ntf); err != nil {
		log.Error(err)
	}

	if err := notify.AddExitContextAction(tray); err != nil {
		log.Error(err)
	}

	ntf.Start()

	if err := tray.SetVisible(true); err != nil {
		log.Fatal(err)
	}
	mainWindow.Run()
}
