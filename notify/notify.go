package notify

import (
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	log "github.com/sirupsen/logrus"
	"time"
)

type Startable interface {
	start()
}

type Stoppable interface {
	stop()
}

func newNotifier(t time.Duration, f func()) *notifier {
	return &notifier{
		runFunc:    f,
		tickerTime: t,
	}
}

type notifier struct {
	runFunc     func()
	tickerTime  time.Duration
	ticker      *time.Ticker
	closeTicker chan bool
}

func (n *notifier) start() {
	log.Info("Starting execution")

	n.closeTicker = make(chan bool)
	n.ticker = time.NewTicker(n.tickerTime)
	go func() {
		for {
			select {
			case <-n.closeTicker:
				return
			case <-n.ticker.C:
				n.runFunc()
			}
		}
	}()
}

func (n *notifier) stop() {
	log.Info("Stopping execution")

	if n.ticker != nil {
		n.ticker.Stop()
		n.closeTicker <- true
	}
}

func OnReady() {
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
			n.start()
		} else {
			item.Check()
			n.stop()
		}
	}
}

func triggerNotification() {
	err := beeep.Notify("Drink Water Notification",
		"An hour has passed, you need to drink some water!", "")
	if err != nil {
		log.Error(err)
	}
}

func handleExit(item *systray.MenuItem) {
	<-item.ClickedCh
	systray.Quit()
}