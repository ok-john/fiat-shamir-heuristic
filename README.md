# fiat-shamir-heuristic

Go implementation of non-interactive proofs of knowledge using the fiat-shamir heuristic.

## Install

```bash
go get github.com/ok-john/fiat-shamir-heuristic
```

## Usage

```go
package main

import (
	"math/big"

	fsh "github.com/ok-john/fiat-shamir-heuristic"
)

func main() {

	// the value you are generating a non-interactive proof of knowledge of.
	message := (&big.Int{}).SetBytes([]byte("you look beautiful today."))

	// the order of the underlying cyclic group.
	order := fsh.Secp256k1Order()

	// Generates the proof.
	proof, err := fsh.Proove(message, order)
	if err != nil {
		panic(err)
	}

	// Verifies the proof.
	if !proof.Verify() {
		panic("failed to verify proof")
	}
}
```
