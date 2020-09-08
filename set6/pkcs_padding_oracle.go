package set6

import (
	"github.com/dbalan/cryptopals/rsa"
	"github.com/dbalan/cryptopals/common"
	"math/big"
	"fmt"
)


const pt = "kick it, CC"
const size = 256

func PKCSPad(msg []byte, n int) []byte {
	pad := []byte{0, 2}
	ps := n - 3 - len(msg)
	padding, err := common.RandBytes(ps)
	if err != nil {
		panic(err)
	}
	
	pad = append(pad,padding...)
	pad = append(pad, 0)
	pad = append(pad, msg...)
	return pad
}

func PKCSOracle() (*big.Int, *big.Int, *big.Int, func(*big.Int) bool) {
	pub, priv, n, err := rsa.GenKeyPair(size)
	if err != nil {
		panic(err)
	}
	m := []byte(pt)
	bytesize := (2*size)/8
	paddedMsg := PKCSPad(m, bytesize)

	nlim := new(big.Int).Exp(big.NewInt(2), big.NewInt(504), nil)
	blim := new(big.Int).Exp(big.NewInt(2), big.NewInt(512), nil)
	// 
	if n.Cmp(nlim) < 0 || n.Cmp(blim) >= 0 {
		panic("n out of bounds")
	}
	
	actualCT := rsa.Encrypt(paddedMsg, pub, n)

	oracle := func(ct *big.Int) bool {
		c := new(big.Int).Set(ct)
		eb := rsa.DecryptRaw(c, priv, n)
		if len(eb) == bytesize-1 {
			// check if eb is PKCS conforming
			return (eb[0] == 2)
		} else {
			return (eb[0] == 0 && eb[1] == 2)
		}
		return false
	}
	return actualCT, pub, n, oracle
}

type Ivl struct {
	A *big.Int
	B *big.Int
}

func testConforming(s, ct, n, pub *big.Int, oracle func(*big.Int) bool) bool {
	c := new(big.Int).Exp(s, pub, n)
	c.Mul(ct, c)
	c.Mod(c, n)
	return oracle(c)
}

func PKCSOracleAttack(ct, pub, n *big.Int, oracle func(*big.Int) bool) {
	one := big.NewInt(1)
	two := big.NewInt(2)

	// specific constants
	k := big.NewInt(512/8)
	n_1 := new(big.Int).Sub(n, one)
	k28 := new(big.Int).Sub(k, two)
	k28.Mul(k28, big.NewInt(8))
	
	B := new(big.Int).Exp(two, k28, nil)
	B2 := new(big.Int).Mul(B, two)
	B3 := new(big.Int).Mul(B, big.NewInt(3))
	B3_1 := new(big.Int).Sub(B3, one)

	
	// Decrypt steps
	
	// 1. Blinding
	// C already is PKCS Conforming, set S0 = 1
	if !oracle(ct) {
		panic("c is not pkcs conforming")
	}
	
	s := new(big.Int).Div(n, B3)
	// 
	M := []Ivl{
		Ivl{B2, B3_1},
	}
	
	i := 1

	for {
		if len(M) == 0 {
			panic("lol what")
		}
		if ! ((i > 1) && (len(M) == 1)) {
			if (i > 1) && (len(M) > 1) {
				fmt.Println("Step 2b: ", M)
				s.Add(s, one)
				
			} else {
				fmt.Println("Step 2a:", M)
			}
			
			for !testConforming(s, ct, n, pub, oracle) {
				s.Add(s, one)
			}

		} else {
			inter := M[0]

			if inter.A.Cmp(inter.B) == 0 {
				// step 4
				fmt.Printf("Found M: %x\n", inter.A)
				return
			}
			fmt.Println("Step 2c: ", M)
			
			r := new(big.Int).Mul(inter.B, s)
			r.Sub(r, B2)
			r.Mul(r, two)
			r.Add(r, n_1)
			r.Div(r, n)

			// incrememnt r until we find an s
			step2c: for {
				// find s
				s = new(big.Int).Mul(r, n)
				s.Add(s, B2)
				s.Add(s, inter.B)
				s.Sub(s, one)
				s.Div(s, inter.B)

				slim := new(big.Int).Mul(r, n)
				slim.Add(slim, B3)
				slim.Add(slim, inter.A)
				slim.Sub(slim, one)
				slim.Div(slim, inter.A)

				for s.Cmp(slim) < 0 {
					if testConforming(s, ct, n, pub, oracle) {
						break step2c
					}
					s.Add(s, one)
				}

				r.Add(r, one)
			}
				
			
		}

		fmt.Println("Step 3: ", M)

		// Step 3
		MNew := []Ivl{}
		for _, inter := range M {
			// ceil division from floor
			// x + n-1/n <= x/n
			r := new(big.Int).Mul(inter.A, s)
			r.Sub(r, B3)
			r.Add(r, n)
			r.Div(r, n)

			rlim := new(big.Int).Mul(inter.B, s)
			rlim.Sub(rlim, B2)
			rlim.Div(rlim, n)

			fmt.Println("R", r, s, inter)
			fmt.Println("rlim", rlim)
			
			for r.Cmp(rlim) <= 0 {
				// new a
				anew := new(big.Int).Mul(r, n)
				anew.Add(anew, B2)
				anew.Add(anew, n)
				anew.Sub(anew, one)
				anew.Div(anew, s)

				if inter.A.Cmp(anew) > 0 {
					anew = inter.A
				}

				// new b
				bnew := new(big.Int).Mul(r, n)
				bnew.Add(anew, B3_1)
				bnew.Div(bnew, s)

				if inter.B.Cmp(bnew) < 0 {
					bnew = inter.B
				}

				MNew = append(MNew, Ivl{anew, bnew})
				r.Add(r, one)
			}
			
		}
		M = []Ivl{}
		M = append(M, MNew...)
		fmt.Println("NEW M", len(M), M)
	}
}
