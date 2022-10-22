package fiat_shamir

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const (
	// The hex encoded order of secp256k1.
	__secp256k1_order_hex = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"
)

// A non-interactive fiat shamir proof of knowledge.
type FiatShamirProof struct {
	t, r, c, g, y, lambda *big.Int
}

// Verifies the proof.
func (proof *FiatShamirProof) Verify() bool {
	l := (&big.Int{}).Exp(proof.g, proof.r, proof.lambda)
	r := (&big.Int{}).Exp(proof.y, proof.c, proof.lambda)
	t := (&big.Int{}).Mul(l, r)
	return t.Cmp(proof.t) == 0

}

// Calculates and returns a non-interactive proof of knowledge
// for `x` using the fiat shamir heuristic.
// g is the group order, x is a the value for which the proof
// of knowledge is generated for.
func Proove(x, g *big.Int) (*FiatShamirProof, error) {
	ret := empty(g)
	ret.lambda = (&big.Int{}).Sub(g, big.NewInt(1))
	ret.y = (&big.Int{}).Exp(g, x, ret.lambda)
	v, err := RandomPointOfGroup(g)
	if err != nil {
		return ret, err
	}
	ret.t.Exp(g, v, ret.lambda)

	c := sha256.Sum256(append(append(ret.g.Bytes(), ret.y.Bytes()...), ret.t.Bytes()...))
	ret.c.SetBytes(c[:])

	ret.r.Mul(ret.c, x)
	ret.r.Sub(v, ret.r)
	ret.r.Mod(ret.r, ret.lambda)

	return ret, nil
}

// Returns the order of curve secp256k1.
func Secp256k1Order() (ret *big.Int) {
	ret, _ = big.NewInt(0).SetString(__secp256k1_order_hex, 16)
	return
}

// Returns some random integer in the ring of integers modulo g.
func RandomPointOfGroup(g *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, g)
}

// Prints the fiat shamir proof in a provided base.
func (p *FiatShamirProof) Print(base int) {
	fmt.Printf("t: %s\n", p.t.Text(base))
	fmt.Printf("r: %s\n", p.r.Text(base))
	fmt.Printf("c: %s\n", p.c.Text(base))
	fmt.Printf("g: %s\n", p.g.Text(base))
	fmt.Printf("y: %s\n", p.y.Text(base))
}

func empty(g *big.Int) *FiatShamirProof {
	return &FiatShamirProof{
		t: big.NewInt(1),
		r: big.NewInt(1),
		c: big.NewInt(1),
		g: g,
		y: big.NewInt(1),
	}
}
