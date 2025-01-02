package context_test

import (
	"context"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	t.Run("`context cancelation` can stop multiple running go routine", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		v := 0

		go func() {
			for {
				select {
				case <-ctx.Done():
					return

				case <-time.After(50 * time.Millisecond):
					v += 1
				}
			}
		}()

		go func() {
			for {
				select {
				case <-ctx.Done():
					return

				case <-time.After(130 * time.Millisecond):
					v += 1
				}
			}
		}()

		time.Sleep(170 * time.Millisecond)
		cancel()
		time.Sleep(300 * time.Millisecond)

		expected := 4
		if v != expected {
			t.Errorf("v:%v , expected:%v", v, expected)
		}
	})
}
