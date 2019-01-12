package set1

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestDecryptAES128ECB(t *testing.T) {
	body, err := ioutil.ReadFile("./7.txt")
	assert.Nil(t, err)

	body = bytes.Replace(body, []byte("\n"), []byte(""), -1)
	decoded := base64decode(body)

	key := "YELLOW SUBMARINE"
	pt, err := DecryptAES128ECB(decoded, []byte(key))
	assert.Nil(t, err)
	fmt.Println("CH7\n", string(pt))
}
