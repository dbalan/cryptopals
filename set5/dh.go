package set5

import (
	"math/big"
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

func smModExp(a, b, p int64) int64 {
	// wrap them in big int since they might actually overflow!
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	bigP := big.NewInt(p)
	bigA.Exp(bigA, bigB, bigP)
	return bigA.Int64()
}

func simpleDHCheck() {
	rand.Seed(time.Now().UnixNano())
	var p int64 = 37
	var g int64 = 5

	ra := rand.Int63()
	rb := rand.Int63()

	a := ra % p
	b := rb % p

	A := smModExp(g, a, p)
	B := smModExp(g, b, p)

	sA := smModExp(B, a, p)
	sB := smModExp(A, b, p)

	if sA != sB {
		panic("Session keys don't match")
	}
}

func primes() (p, g *big.Int) {
	p = &big.Int{}
	g = &big.Int{}
	if _, ok := p.SetString(strP, 16); !ok {
		panic("error: noparse.strP")
	}

	if _, ok := g.SetString(strG, 16); !ok {
		panic("error: noparse.strG")
	}
	return
}

func keypair(p, g *big.Int) (pub, priv *big.Int) {
	priv = big.NewInt(rand.Int63())
	priv.Mod(priv, p)

	pub = &big.Int{}
	pub.Exp(g, priv, p)
	return pub, priv
}

func sessionKey(pubTheir, privMine, p *big.Int) *big.Int {
	skey := &big.Int{}
	skey.Exp(pubTheir, privMine, p)
	return skey
}

func NISTDHCheck() {
	p, g := primes()

	// personA
	A, a := keypair(p, g)
	// personB
	B, b := keypair(p, g)

	// transfer
	// personA
	skeyA := sessionKey(B, a, p)

	// personB
	skeyB := sessionKey(A, b, p)

	if skeyA.Cmp(skeyB) != 0 {
		panic("session keys don't match")
	}
}
