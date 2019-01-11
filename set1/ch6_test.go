package set1

import (
	"encoding/base64"
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
		data += "ab"
	}

	result := sliceAndDice([]byte(data), 2)

	expected := [][]byte{}
	expected = append(expected, []byte("aaaaaaaaaa"))
	expected = append(expected, []byte("bbbbbbbbbb"))

	assert.Equal(t, expected, result)
}

func TestFindKey(t *testing.T) {
	cipherText := decodeHexString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	keySize := findKeySize(cipherText)

	assert.Equal(t, 3, keySize)
	key := findKey(cipherText, keySize)
	assert.Equal(t, "ICE", string(key))
}

func TestCH6FindKey(t *testing.T) {
	body, err := ioutil.ReadFile("./6.txt")
	assert.Nil(t, err)

	encoded := string(body)
	encoded = strings.Replace(encoded, "\n", "", -1)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	assert.Nil(t, err)

	keySize := findKeySize(decoded)
	key := findKey(decoded, keySize)
	//	decrypted := decryptRepeatXOR(decoded, key)
	fmt.Printf("KEYSIZE: %d \n KEY: %s\n", keySize, string(key))

}
