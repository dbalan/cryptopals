package set4

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"net/http"
	"time"
)

const startingPoint = "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"

// const startingPoint = "587b2c47492fd3c153aaa23e2504cf2b397zzzzz"

func makeRequest2(file, signature string) (bool, time.Duration) {
	u := fmt.Sprintf("http://localhost:9000/test?file=%s&signature=%s",
		file, signature)
	n := time.Now()
	resp, err := http.Get(u)
	if err != nil {
		panic("got error talking to server: " + err.Error())
	}

	resp.Body.Close()
	status := resp.StatusCode == 200
	ts := time.Since(n)
	return status, ts
}

func insecComp(fst, snd []byte) bool {
	l := len(fst)
	if l != len(snd) {
		return false
	}

	for i := 0; i < l; i++ {
		if fst[i] != snd[i] {
			return false
		}
		time.Sleep(20 * time.Millisecond)
	}

	return true

}

func checkSig(file, sig string) bool {
	actualsig := common.EncodeHexString(
		SHA1MAC([]byte("$upEr$ECr4TK$y"), []byte(file)),
	)
	return insecComp([]byte(actualsig), []byte(sig))
}

func makeRequest(file, sig string) (bool, time.Duration) {
	n := time.Now()
	flag := checkSig(file, sig)
	ts := time.Since(n)
	return flag, ts
}

func bruteforceSig(file string) {
	// find a proper signature for file
	found := []byte(startingPoint)

	var tprev time.Duration = 0
outer:
	for l := 0; l < len(found); l++ {
	try_again:
		var highest time.Duration
		var probable byte
		for p := 0; p <= 0xf; p++ {
			chr := common.EncodeHex(byte(p))

			found[l] = chr
			suc, t := makeRequest(file, string(found))
			if suc {
				fmt.Printf("Singature found! file=%s sig=%s\n", file, string(found))
				break outer
			}
			if t.Nanoseconds() > highest.Nanoseconds() {
				highest = t
				probable = chr
			}

		}

		if highest.Nanoseconds() < tprev.Nanoseconds() || (highest.Seconds()-tprev.Seconds()) < (18*time.Millisecond).Seconds() {
			fmt.Printf("prev: %f now: %f - l at: %d\n", tprev.Seconds(), highest.Seconds(), l)
			// l = l -
			tprev = 0
			goto try_again
		}

		tprev = highest

		found[l] = probable
		fmt.Println("prefix: ", string(found[0:l]))
	}

}
