package set2

import (
	"fmt"
	"strings"
)

type KVMap struct {
	s    map[string]string
	keys []string
}

func decodeKV(in string) *KVMap {
	pairs := strings.Split(in, "&")
	decoded := make(map[string]string)
	keys := []string{}

	for _, p := range pairs {
		splt := strings.Split(p, "=")
		if len(splt) != 2 {
			// discard
			continue
		}
		keys = append(keys, splt[0])
		decoded[splt[0]] = splt[1]
	}

	return &KVMap{decoded, keys}

}

func (kv *KVMap) encode() string {
	acc := []string{}

	for _, k := range kv.keys {
		acc = append(acc, fmt.Sprintf("%s=%s", k, kv.s[k]))
	}

	return strings.Join(acc, "&")
}

func profileFor(email string) *KVMap {
	email = strings.Replace(email, "&", "", -1)
	email = strings.Replace(email, "=", "", -1)

	return decodeKV(fmt.Sprintf("email=%s&uid=10&role=user", email))
}
