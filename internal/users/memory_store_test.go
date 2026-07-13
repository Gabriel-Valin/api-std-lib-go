package users

import (
	"sync"
	"testing"
)

func TestMemoryStore_ConcurrentList(t *testing.T) {
	store := NewMemoryStore()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			_ = store.List()
		}()
	}

	wg.Wait()
}
