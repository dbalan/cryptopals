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
		wg.Add(1)
		go func() {
			defer wg.Done()
			simpleDHCheck()
		}()
	}
	wg.Wait()
}

func TestPrimes(t *testing.T) {
	p, g := primes()
	assert.Equal(t, strP, p.Text(16))
	assert.Equal(t, strG, g.Text(16))
}

func TestNISTDHCheck(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			NISTDHCheck()
		}()
	}
	wg.Wait()
}
