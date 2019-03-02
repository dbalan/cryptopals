package set6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	msg := []byte("It seems")
	s, err := NewSvr()
	assert.Nil(t, err)
	ct, pub, n := s.Encrypt(msg)

	_, err = s.Decrypt(ct)
	assert.Nil(t, err)

	_, err = s.Decrypt(ct)
	assert.NotNil(t, err)

	nct, multi := tamperRsa(ct, pub, n)

	m, err := s.Decrypt(nct)
	assert.Nil(t, err)

	recov, err := recoverRsa(m, multi, n)
	assert.Nil(t, err)

	assert.Equal(t, msg, recov)
}
