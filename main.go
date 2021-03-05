package main

import (
	"drinkwater-go/notify"
	"github.com/getlantern/systray"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	initLogging()
	systray.Run(func() { notify.OnReady() }, func() { notify.OnExit() })
}

func initLogging() {
	file, err := os.OpenFile("drinkwater.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
