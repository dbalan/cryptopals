package set5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegularLogin(t *testing.T) {
	N, _ = primes()
	pass := randPasswd()
	server := NewServer(pass)
	assert.Equal(t, false, login(server, "hello world"))
	assert.Equal(t, true, login(server, pass))
}

func TestLoginZero(t *testing.T) {
	N, _ = primes()
	server := NewServer(randPasswd())
	assert.Equal(t, true, loginWithZero(server))
}

func TestLoginWithN(t *testing.T) {
	N, _ = primes()
	server := NewServer(randPasswd())
	assert.Equal(t, true, loginWithN(server))
}
