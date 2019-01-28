package set1

import (
	"github.com/dbalan/cryptopals/common"
	"io/ioutil"
	"strings"
)

func detectCipherText(in []string) string {
	low := float64(0xffffffff)
	cipher := ""
	for _, pt := range in {
		cur := common.TextScore([]byte(pt))
		if cur < low {
			low = cur
			cipher = pt
		}
	}
	return cipher
}

func readFile(path string) ([]string, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(body), "\n"), nil
}
