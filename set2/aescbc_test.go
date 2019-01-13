package set2

import (
	"bytes"
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestAESCBC(t *testing.T) {
	// Twenty Thousand Leagues under the Sea :-)
	pt := "On my arrival at New York the question was at its height.  The theory of the floating island, and the unapproachable sandbank, supported by minds little competent to form a judgment, was abandoned.  And, indeed, unless this shoal had a machine in its stomach, how could it change its position with such astonishing ra..."

	blocksize := 16
	key := "YELLOW SUBMARINE"
	iv := []byte{}
	for i := 0; i < blocksize; i++ {
		iv = append(iv, byte(0))
	}

	enc, err := EncAES128CBC([]byte(pt), iv, []byte(key))
	assert.Nil(t, err)
	assert.Equal(t, len(pt), len(enc))

	dec, err := DecAES128CBC(enc, iv, []byte(key))
	assert.Nil(t, err)
	assert.Equal(t, pt, string(dec))
}

// find solution to ch10
func TestCh10(t *testing.T) {
	body, err := ioutil.ReadFile("./10.txt")
	assert.Nil(t, err)

	// remove new lines
	body = bytes.Replace(body, []byte("\n"), []byte(""), -1)

	decoded := common.DecodeB64(body)
	key := []byte("YELLOW SUBMARINE")

	iv := []byte{}
	for i := 0; i < 16; i++ {
		iv = append(iv, byte(0))
	}

	dec, err := DecAES128CBC(decoded, iv, []byte(key))
	assert.Nil(t, err)
	fmt.Println(string(dec))
}
