package set5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCipherTextE3(t *testing.T) {
	_, _, err := getCipherTextE3([]byte("hey"))
	assert.Nil(t, err)
}

func TestHastadBCAttack(t *testing.T) {
	err := HastadBCAttack()
	assert.Nil(t, err)
}
