package set4

import (
	"errors"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set2"
	"strings"
)

var (
	HighAsciiErr = errors.New("high ascii errors")
)

func keyAsIVOracle(in string, key []byte) ([]byte, error) {
	in = strings.Replace(in, ";", "%3B", -1)
	in = strings.Replace(in, "=", "%3D", -1)
	pt := []byte("comment1=cooking%20MCs;userdata=" + in + ";comment2=%20like%20a%20pound%20of%20bacon")
	lpt := len(pt)
	if (lpt % 16) != 0 {
		newlen := lpt + 16 - (lpt % 16)
		padded, err := set2.PKCS7Padding(pt, newlen)
		if err != nil {
			return nil, err
		}
		pt = padded
	}
	enc, err := set2.EncAES128CBC(pt, key, key)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func modifyCT(ct []byte) []byte {
	blocks := common.Blocks(ct, 16)
	if len(blocks) < 3 {
		panic("not enough!")
	}

	newct := blocks[0]
	newct = append(newct, common.Repeat(16, byte(0))...)
	newct = append(newct, blocks[0]...)
	return newct
}

func userDecrypt(ct []byte, key []byte) ([]byte, error) {
	dec, _ := set2.DecAES128CBC(ct, key, key)

	for _, d := range dec {
		if int(d) > 0x7f {
			return dec, HighAsciiErr
		}
	}
	return nil, nil
}

func AttackKeyAsIV() error {
	key, err := common.RandBytes(16)
	if err != nil {
		return err
	}

	ct, err := keyAsIVOracle("something", key)
	if err != nil {
		return err
	}

	// attacker
	newct := modifyCT(ct)

	// user
	dec, err := userDecrypt(newct, key)
	if err == nil || err != HighAsciiErr {
		return errors.New("we havve a problem")
	}

	// attacker again
	blocks := common.Blocks(dec, 16)
	recoverd, err := common.XORBlk(blocks[0], blocks[2])
	if err != nil {
		return err
	}

	if !common.EqualBlocks(recoverd, key) {
		return errors.New("wrong key")
	}
	return nil
}
