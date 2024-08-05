package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		count := DataMutex{}
		count.Counter()
		count.Counter()
		count.Counter()
		if count.value != 3 {
			t.Errorf("got %d want %d", count.value, 3)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		count := DataMutex{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				count.Counter()
				wg.Done()
			}()
		}
		wg.Wait()
		if count.value != wantedCount {
			t.Errorf("got %d want %d", count.value, wantedCount)
		}
	})
}
func TestEven(t *testing.T) {
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		count := DataMutex{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				count.Even()
				count.Odd()
				wg.Done()
			}()
		}
		wg.Wait()
		if count.value != 21 {
			t.Errorf("got %d want %d", count.value, 21)
		}
	})
}
