package channel_test

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	t.Run("`channel + select` can tell a running goroutine to stop", func(t *testing.T) {
		t.Parallel()

		stop := make(chan bool)
		v := 0

		go func() {
			for {
				select {
				case <-stop:
					return

				case <-time.After(50 * time.Millisecond):
					v += 1
				}
			}
		}()
		time.Sleep(170 * time.Millisecond)
		stop <- true
		time.Sleep(300 * time.Millisecond)

		expected := 3
		if v != expected {
			t.Errorf("v:%v , expected:%v", v, expected)
		}
	})

	t.Run("`channel` inform one listener a time", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)
		v := 0

		go func() {
			v += <-ch
		}()

		go func() {
			v += <-ch
		}()

		ch <- 1
		time.Sleep(50 * time.Millisecond)

		expected := 1
		if v != expected {
			t.Errorf("v:%v , expected:%v", v, expected)
		}
	})
	t.Run("`close(channel)` can stop multiple running go routine", func(t *testing.T) {
		t.Parallel()

		stop := make(chan struct{})
		v := 0

		go func() {
			for {
				select {
				case <-stop:
					return

				case <-time.After(50 * time.Millisecond):
					v += 1
				}
			}
		}()

		go func() {
			for {
				select {
				case <-stop:
					return

				case <-time.After(130 * time.Millisecond):
					v += 1
				}
			}
		}()

		time.Sleep(170 * time.Millisecond)
		close(stop)
		time.Sleep(300 * time.Millisecond)

		expected := 4
		if v != expected {
			t.Errorf("v:%v , expected:%v", v, expected)
		}
	})
}
