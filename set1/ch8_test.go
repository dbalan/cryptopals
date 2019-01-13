package set1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestDetectAES128ECB(t *testing.T) {
	body, err := ioutil.ReadFile("./8.txt")
	assert.Nil(t, err)
	data := string(body)

	for _, ct := range strings.Split(data, "\n") {
		times := repeatingBlocks([]byte(ct))
		if times > 0 {
			fmt.Printf("REPEATING BLOCKS in (%d times): %s", times, ct)
		}
	}

}
