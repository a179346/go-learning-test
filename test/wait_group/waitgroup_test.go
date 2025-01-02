package waitgroup_test

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	t.Run("WaitGroup can be used to wait until multiple tasks to finish", func(t *testing.T) {
		t.Parallel()

		var wg sync.WaitGroup
		wg.Add(2)
		v := 0

		go func() {
			time.Sleep(1 * time.Second)
			v += 1
			wg.Done()
		}()

		go func() {
			time.Sleep(2 * time.Second)
			v += 2
			wg.Done()
		}()

		wg.Wait()
		expected := 3
		if v != expected {
			t.Errorf("v is %v. Expected: %v", v, expected)
		}
	})
}
