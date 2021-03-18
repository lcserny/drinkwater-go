package notify

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"strings"
	"testing"
	"time"
)

func TestNotifier(t *testing.T) {
	t.Run("notifier can be started to run given function", func(t *testing.T) {
		functionRan := make(chan struct{})

		notifier := NewNotifier(5*time.Millisecond, func() { <-functionRan })

		select {
		case <-functionRan:
			t.Error("did not expect function to be run yet")
		case <-time.After(10 * time.Millisecond):
			return
		}

		notifier.Start()

		select {
		case <-functionRan:
			return
		case <-time.After(10 * time.Millisecond):
			t.Error("expected function to be run")
		}

		close(functionRan)
	})

	t.Run("starting notifier logs message", func(t *testing.T) {
		assertLogContains(t, NewNotifier(1*time.Hour, func() {}).Start, startedExecution)
	})

	t.Run("started notifier can be paused", func(t *testing.T) {
		functionRan := make(chan struct{})

		notifier := NewNotifier(5*time.Millisecond, func() { <-functionRan })
		notifier.Start()
		notifier.Pause()

		select {
		case <-functionRan:
			t.Error("did not expect paused function to be run")
		case <-time.After(10 * time.Millisecond):
			return
		}

		close(functionRan)
	})

	t.Run("pausing notifier logs message", func(t *testing.T) {
		assertLogContains(t, NewNotifier(1*time.Hour, func() {}).Pause, pausedExecution)
	})

	t.Run("paused notifier can be resumed", func(t *testing.T) {
		functionRan := make(chan struct{})

		notifier := NewNotifier(5*time.Millisecond, func() { <-functionRan })
		notifier.Start()
		notifier.Pause()

		select {
		case <-functionRan:
			t.Error("did not expect paused function to be run")
		case <-time.After(10 * time.Millisecond):
			return
		}

		notifier.Resume()

		select {
		case <-functionRan:
			return
		case <-time.After(10 * time.Millisecond):
			t.Error("expected resumed function to be run")
		}

		close(functionRan)
	})

	t.Run("resuming notifier logs message", func(t *testing.T) {
		assertLogContains(t, NewNotifier(1*time.Hour, func() {}).Resume, resumedExecution)
	})
}

func assertLogContains(t testing.TB, notifierFunc func(), wantedLog string) {
	t.Helper()

	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	notifierFunc()

	if !strings.Contains(buffer.String(), wantedLog) {
		t.Errorf("expected to log message %q", wantedLog)
	}
}
