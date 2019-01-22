package set2

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptEmailRole(t *testing.T) {
	key, err := common.RandBytes(16)
	assert.Nil(t, err)

	oracle := func(e string) []byte {
		return encryptEmail(e, key)
	}

	// pad role to a byte
	paddedRole, err := PKCS7Padding([]byte("admin"), 16)
	assert.Nil(t, err)

	// prefix it with enough to move it to a new block
	prefix := []byte{}
	for i := 0; i < 16-len("email="); i++ {
		prefix = append(prefix, byte('f'))
	}

	email := append(prefix, paddedRole...)
	ct := oracle(string(email))

	// forged ct block -> would return admin when decrypted
	forged := ct[16:32]

	// email of length that moves the role name to its own block
	actualEmail := "admin@what.co"
	newct := oracle(actualEmail)

	// strip actual last block and add ours
	newct = newct[0 : len(newct)-16]
	newct = append(newct, forged...)

	PT := decryptKV(newct, key)
	fmt.Println("forged plain text: ", PT)
}
