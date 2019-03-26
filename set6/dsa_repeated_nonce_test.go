package set6

import (
	"fmt"
	"github.com/dbalan/cryptopals/sha"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadData(t *testing.T) {
	pairs, err := readData()

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(pairs))
}

func TestFindRepeating(t *testing.T) {
	p1, p2, err := findRepeatingK()
	assert.Nil(t, err)
	if p1.R.Cmp(p2.R) != 0 {
		t.Error("fail")
	}
}

func TestFindPrivateKey(t *testing.T) {
	priv := findPrivateKey()

	privHex := priv.Text(16)
	got := fmt.Sprintf("%x", sha.SHA([]byte(privHex)))

	expected := "ca8f6f7c66fa362d40760d135b763eb8527d3d52"
	assert.Equal(t, expected, got)
}
