package notify

import (
	log "github.com/sirupsen/logrus"
	"time"
)

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
