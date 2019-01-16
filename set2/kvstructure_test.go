package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileFor(t *testing.T) {
	pmap := profileFor("foo@bar.com")
	emap := pmap.encode()

	assert.Equal(t, "email=foo@bar.com&uid=10&role=user", emap)
}
