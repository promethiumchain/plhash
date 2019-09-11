package main

import "math/big"

// ZeroBigInt returns a new big int with zero value
func ZeroBigInt() *big.Int {
	return big.NewInt(0)
}
