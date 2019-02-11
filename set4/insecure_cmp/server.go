package main

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set4"
	"net/http"
	"time"
)

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
		set4.SHA1MAC([]byte("$upEr$ECr4TK$y"), []byte(file)),
	)
	return insecComp([]byte(actualsig), []byte(sig))
}

func attackTarget(w http.ResponseWriter, r *http.Request) {
	keys := []string{"file", "signature"}
	r.ParseForm()
	for _, k := range keys {
		if _, ok := r.Form[k]; !ok {
			http.Error(w, "Set "+k, http.StatusBadRequest)
			return
		}
	}

	// hardcoded ugliness..
	file := r.Form["file"][0]
	sig := r.Form["signature"][0]

	if !checkSig(file, sig) {
		http.Error(w, "error: Bad Signature", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "accepted: good signature")

}

func main() {
	http.HandleFunc("/test", attackTarget)
	http.ListenAndServe(":9000", nil)
}
