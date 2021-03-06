package dsa

import (
	"crypto/rand"
	"github.com/dbalan/cryptopals/sha"
	"math/big"
)

const (
	defaultP = "800000000000000089e1855218a0e7dac38136ffafa72eda7" +
		"859f2171e25e65eac698c1702578b07dc2a1076da241c76c6" +
		"2d374d8389ea5aeffd3226a0530cc565f3bf6b50929139ebe" +
		"ac04f48c3c84afb796d61e5a4f9a8fda812ab59494232c7d2" +
		"b4deb50aa18ee9e132bfa85ac4374d7f9091abc3d015efc87" +
		"1a584471bb1"
	defaultQ = "f4f47f05794b256174bba6e9b396a7707e563c5b"
	defaultG = "5958c9d3898b224b12672c0b98e06c60df923cb8bc999d119" +
		"458fef538b8fa4046c8db53039db620c094c9fa077ef389b5" +
		"322a559946a71903f990f1f7e0e025e2d7f7cf494aff1a047" +
		"0f5b64c36b625a097f1651fe775323556fe00b3608c887892" +
		"878480e99041be601a62166ca6894bdd41a7054ec89f756ba" +
		"9fc95302291"
)

type ParamFn func() (*big.Int, *big.Int, *big.Int)

func GetDefaultParams() (p, q, g *big.Int) {
	p = new(big.Int)
	p.SetString(defaultP, 16)

	q = new(big.Int)
	q.SetString(defaultQ, 16)

	g = new(big.Int)
	g.SetString(defaultG, 16)
	return
}

// x = priv
// y = pub
func KeyPair(pf ParamFn) (x, y *big.Int, err error) {
	p, q, g := pf()

	x, err = rand.Int(rand.Reader, q)
	if err != nil {
		return
	}

	y = new(big.Int).Exp(g, x, p)
	return
}

func Sign(msg []byte, x *big.Int, pf ParamFn) (r, s *big.Int, err error) {
	r, s, _, err = signInternal(msg, x, pf)
	return
}

// Sign with SHA-1
func signInternal(msg []byte, x *big.Int, pf ParamFn) (r, s, k *big.Int, err error) {
	p, q, g := pf()
	hs := hsmsg(msg)

newk:
	k, err = rand.Int(rand.Reader, q)
	if err != nil {
		return
	}

	r = new(big.Int).Exp(g, k, p)
	r.Mod(r, q)

	if r.Cmp(big.NewInt(0)) == 0 {
		// try with new k
		goto newk
	}

	xr := new(big.Int).Mul(x, r)
	hs.Add(hs, xr)

	s = new(big.Int).Exp(k, big.NewInt(-1), q)
	s.Mul(s, hs)
	s.Mod(s, q)

	if s.Cmp(big.NewInt(0)) == 0 {
		// try with new k
		goto newk
	}
	return
}

func Verify(msg []byte, r, s, y *big.Int, pf ParamFn) bool {
	p, q, g := pf()
	zero := big.NewInt(0)

	if r.Cmp(zero) <= 0 || r.Cmp(q) >= 0 {
		return false
	}

	if s.Cmp(zero) <= 0 || s.Cmp(q) >= 0 {
		return false
	}

	w := new(big.Int).Exp(s, big.NewInt(-1), q)
	hs := hsmsg(msg)

	u1 := new(big.Int).Mul(hs, w)
	u1.Mod(u1, q)

	u2 := new(big.Int).Mul(r, w)
	u2.Mod(u2, q)

	v1 := new(big.Int).Exp(g, u1, p)
	v2 := new(big.Int).Exp(y, u2, p)
	v := new(big.Int).Mul(v1, v2)
	v.Mod(v, p)
	v.Mod(v, q)

	return v.Cmp(r) == 0
}

/*
 *         (s * k) - H(msg)
 *     x = ----------------  mod q
 *                 r
 */
func ComputeKey(k, r, s, hs *big.Int) *big.Int {
	_, q, _ := GetDefaultParams()
	/*
		println("k = 0x" + k.Text(16))
		println("r = 0x" + r.Text(16))
		println("s = 0x" + s.Text(16))
		println("hs = 0x" + hs.Text(16))
		println("q  = 0x" + q.Text(16))
	*/
	u := new(big.Int).Set(s)
	u.Mul(u, k)
	u.Sub(u, hs)

	rinv := big.NewInt(-1)
	rinv.Exp(r, rinv, q)
	u.Mul(u, rinv)
	u.Mod(u, q)
	return u
}

func hsmsg(msg []byte) *big.Int {
	return new(big.Int).SetBytes(sha.SHA(msg))
}
