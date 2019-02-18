package set5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvilSSRP(t *testing.T) {

	pass := randPasswd()
	evil := newEvilSSRP()
	assert.Equal(t, false, smplSRPLogin(evil, pass))
	assert.Equal(t, pass, evil.GetPassword())
}
