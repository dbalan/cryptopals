package set3

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestEncodeUint64(t *testing.T) {
	cs := []struct {
		In  uint64
		Out []int
	}{
		{1, []int{1, 0, 0, 0, 0, 0, 0, 0}},
		{17, []int{17, 0, 0, 0, 0, 0, 0, 0}},
		{257, []int{1, 1, 0, 0, 0, 0, 0, 0}},
		{168430090, []int{10, 10, 10, 10, 0, 0, 0, 0}},
	}

	for _, c := range cs {
		out := encodeUint64(c.In)

		expected := []byte{}
		for _, e := range c.Out {
			expected = append(expected, byte(e))
		}
		assert.Equal(t, expected, out)
	}
}

func TestDecAESCTR(t *testing.T) {
	ct := []byte("L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==")
	decoded := common.DecodeB64(ct)

	key := []byte("YELLOW SUBMARINE")
	nonce := uint64(0)

	resp, err := AES128CTR(decoded, key, nonce)
	assert.Nil(t, err)
	fmt.Printf("CH18: %s\n", string(resp))
}
