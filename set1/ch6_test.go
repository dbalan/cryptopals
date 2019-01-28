package set1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	dist := hammingDistance(
		[]byte("this is a test"), []byte("wokka wokka!!!"),
	)
	assert.Equal(t, 37, dist)
}

func TestFindKeySize(t *testing.T) {
	cipherText := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	size := findKeySize(decodeHexString(cipherText))
	assert.Equal(t, 3, size)

}

func TestSliceAndDice(t *testing.T) {
	data := ""
	for i := 0; i < 10; i++ {
		data += "abcd"
	}

	result := sliceAndDice([]byte(data), 2)

	expected := [][]byte{}
	expected = append(expected, []byte("acacacacacacacacacac"))
	expected = append(expected, []byte("bdbdbdbdbdbdbdbdbdbd"))

	assert.Equal(t, expected, result)
}

func TestCh6(t *testing.T) {
	body, err := ioutil.ReadFile("./6.txt")
	assert.Nil(t, err)
	data := string(body)
	data = strings.Replace(data, "\n", "", -1)
	decoded := base64decode([]byte(data))

	size := findKeySize(decoded)
	fmt.Println("KEYSIZE: ", size)

	key := FindKey(decoded, 29)
	fmt.Println("KEY", key)

	decrypted := DecryptRepeatXOR(decoded, key)
	fmt.Println("partial: ", string(decrypted))
}
