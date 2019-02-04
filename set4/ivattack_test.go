package set4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttackKeyAsIV(t *testing.T) {
	err := AttackKeyAsIV()
	assert.Nil(t, err)

}
