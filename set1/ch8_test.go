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
		if repeatingBlocks([]byte(ct)) {
			fmt.Println("REPEATING BLOCKS: ", ct)
		}
	}

}
