package set5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegularLogin(t *testing.T) {
	assert.Equal(t, false, login("hello world"))
	assert.Equal(t, true, login("5upers4cr4t"))
}
