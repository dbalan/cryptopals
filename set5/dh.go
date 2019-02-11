package set5

import (
	"math"
	"math/rand"
	"time"
)

/* p = 37, g = 5
 * a = rand() mod p
 * A = g^a mod p (pubkey for P1)
 * b = rand() mod p
 * B = g^b mod p (pubkey P2)
 * s = B^a mod p = A^b mod p - session key
 */

const (
	// declare constans
	strP = "ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024" +
		"e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd" +
		"3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec" +
		"6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f" +
		"24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361" +
		"c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552" +
		"bb9ed529077096966d670c354e4abc9804f1746c08ca237327fff" +
		"fffffffffffff"
	strG = "2"
)

func simpleDHCheck() {
	rand.Seed(time.Now().UnixNano())
	p := 37
	g := 5

	a := rand.Int() % p
	b := rand.Int() % p

	A := int(math.Pow(float64(g), float64(a))) % p
	B := int(math.Pow(float64(g), float64(b))) % p

	sA := int(math.Pow(float64(B), float64(a))) % p
	sB := int(math.Pow(float64(A), float64(a))) % p

	if sA != sB {
		panic("Session keys don't match")
	}
}
