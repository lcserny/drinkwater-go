package notify

import (
	"testing"
	"time"
)

// TODO: redo after refactor
func TestNotifier(t *testing.T) {
	t.Run("notifier can be started to run given function", func(t *testing.T) {
		functionRan := false
		notifier := newNotifier(5*time.Millisecond, func() { functionRan = true })

		time.Sleep(10 * time.Millisecond)

		if functionRan {
			t.Error("did not expect function to be run yet")
		}

		go notifier.start()

		time.Sleep(10 * time.Millisecond)

		if !functionRan {
			t.Error("expected function to be run")
		}
	})

	t.Run("started notifier can be paused", func(t *testing.T) {

	})

	t.Run("paused notifier can be resumed", func(t *testing.T) {

	})
}
