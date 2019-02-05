package set4

import (
	"github.com/dbalan/cryptopals/sha"
)

func SHA1MAC(key, msg []byte) (mac []byte) {
	nm := append(key, msg...)
	return sha.SHA(nm)
}
