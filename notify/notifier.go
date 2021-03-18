package notify

import (
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	startedExecution = "Starting execution"
	pausedExecution  = "Pausing execution"
	resumedExecution = "Resuming execution"
)

func NewNotifier(t time.Duration, f func()) *notifier {
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

func (n *notifier) Start() {
	log.Info(startedExecution)
	n.ticker = time.NewTicker(n.tickerTime)
	go func() {
		for {
			<-n.ticker.C
			n.runFunc()
		}
	}()
}

func (n *notifier) Pause() {
	log.Info(pausedExecution)
	if n.ticker != nil {
		n.ticker.Stop()
	}
}

func (n *notifier) Resume() {
	log.Info(resumedExecution)
	if n.ticker != nil {
		n.ticker.Reset(n.tickerTime)
	}
}
