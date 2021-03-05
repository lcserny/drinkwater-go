package main

import (
	"drinkwater-go/notify"
	"github.com/getlantern/systray"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	f := initLogging()
	defer f.Close()

	systray.Run(func() { notify.OnReady() }, func() { notify.OnExit() })
}

func initLogging() *os.File {
	file, err := os.OpenFile("drinkwater.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}
