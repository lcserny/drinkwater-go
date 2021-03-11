package main

import (
	"drinkwater-go/notify"
	"github.com/getlantern/systray"
	log "github.com/sirupsen/logrus"
	"os"
)

const logFile = "drinkwater.log"

func main() {
	f := initLogging()
	defer f.Close()

	systray.Run(notify.OnReady, notify.OnExit)
}

func initLogging() *os.File {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}
