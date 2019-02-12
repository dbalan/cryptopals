package set5

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestModExp(t *testing.T) {
	assert.Equal(t, int64(3), smModExp(2, 26, 37))
	assert.Equal(t, int64(3), smModExp(5, 34, 37))
}

func TestSimpleDH(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			simpleDHCheck()
		}()
	}
	wg.Wait()
}
