package notify

import (
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	title   = "Drink Water Notification"
	message = "An hour has passed, you need to drink some water!"
)

type pauseable interface {
	start()
	pause()
	unpause()
}

func newNotifier(t time.Duration, f func()) *notifier {
	return &notifier{
		runFunc:    f,
		tickerTime: t,
	}
}

type notifier struct {
	runFunc    func()
	tickerTime time.Duration
	ticker     *time.Ticker
}

func (n *notifier) start() {
	log.Info("Starting execution")
	n.ticker = time.NewTicker(n.tickerTime)
	for {
		<-n.ticker.C
		n.runFunc()
	}
}

func (n *notifier) pause() {
	log.Info("Pausing execution")
	if n.ticker != nil {
		n.ticker.Stop()
	}
}

func (n *notifier) unpause() {
	log.Info("Unpausing execution")
	if n.ticker != nil {
		n.ticker.Reset(n.tickerTime)
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
			n.unpause()
		} else {
			item.Check()
			n.pause()
		}
	}
}

func triggerNotification() {
	if err := beeep.Notify(title, message, ""); err != nil {
		log.Error(err)
	}
}

func handleExit(item *systray.MenuItem) {
	<-item.ClickedCh
	systray.Quit()
}
