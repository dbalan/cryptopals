package set1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCh4(t *testing.T) {
	data, err := readFile("./4.txt")
	assert.Nil(t, err)
	fmt.Println(detectCipherText(data))
}
