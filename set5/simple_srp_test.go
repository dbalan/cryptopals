package set5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginOk(t *testing.T) {
	pass := randPasswd()
	svr := newSimpleSRPSvr(pass)
	assert.Equal(t, false, smplSRPLogin(svr, "hello"))
	assert.Equal(t, true, smplSRPLogin(svr, pass))
}
