package set6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTamperDSAParam0(t *testing.T) {
	msgs := []string{"Hello, world", "Goodbye, world"}
	for _, m := range msgs {
		err := tamperP1([]byte(m))
		assert.Nil(t, err)
	}
}
