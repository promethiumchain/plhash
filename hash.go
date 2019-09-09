package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"math/big"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

var mathFuncs *MathFuncs
var targetBits uint64 = 2
var baselinePrecisionDigits uint = 128

// init is called once on startup and creates a new list of math funcs
func init() {
	mathFuncs = NewFuncList()
	PrintMessage("math functions list initialized with length : ", len(mathFuncs.FuncList))
}

// Hash32 takes a hash type and returns the byte slice of the hash
func Hash32(h hash.Hash, data []byte) [32]byte {
	var out [32]byte
	h.Write(data)
	h.Sum(out[:0])
	h.Reset()
	return out
}

// Hash64  takes a hash type and returns the byte slice of the hash
func Hash64(h hash.Hash, data []byte) [64]byte {
	var out [64]byte
	h.Write(data)
	h.Sum(out[:0])
	h.Reset()
	return out
}

// HashPassA represents the SHA256 pass of the algo and returns a big int
// it first hashes the data byte slice with the given algo and then creates
// a 64 byte slice of it by reversing the 32 bytes of the original and appending it
func HashPassA(data []byte, index int) (*big.Int, error) {
	h := sha256.New()
	hash := Hash32(h, data)
	hash64 := ReverseBytes32To64(hash)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassB represents the first SHA3.512 pass of the algo and returns a big int
func HashPassB(data []byte, index int) (*big.Int, error) {
	h := sha3.New512()
	hash64 := Hash64(h, data)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassC represents the second SHA3.512 pass of the algo and returns a big int
func HashPassC(data []byte, index int) (*big.Int, error) {
	h := sha3.New512()
	hash64 := Hash64(h, data)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassD represents the BLAKE2S pass of the algo and returns a big int
func HashPassD(data []byte, index int) (*big.Int, error) {
	h, _ := blake2s.New256(nil)
	hash := Hash32(h, data)
	hash64 := ReverseBytes32To64(hash)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassE represents the BLAKE2B pass of the algo and returns a big int
func HashPassE(data []byte, index int) (*big.Int, error) {
	h, _ := blake2b.New512(nil)
	hash64 := Hash64(h, data)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// calcPass takes a hash and passes it via the math funcs, returns a big int
func calcPass(in []byte, index int) (*big.Int, error) {
	bn := ByteToBigInt(in)
	fx := BigIntToBigFloat(bn, baselinePrecisionDigits)
	fnc := mathFuncs.FuncList[index]
	fn := fnc(fx)
	fs, err := RemoveDecFromFloat(fn)
	if err != nil {
		return nil, err
	}
	fbn, errN := SelectLastDigits(fs, 1*5) // depth is multiply to diff
	if errN != nil {
		return nil, errN
	}
	return fbn, nil
}

// CalcFinalHash returns the final hash of the hash passes
func CalcFinalHash(a, b, c, d, e *big.Int) *big.Int {
	fhp := sha3.New512()
	sumAB := ZeroBigInt()
	mulCD := ZeroBigInt()
	subABCD := ZeroBigInt()
	divE := ZeroBigInt()
	sumAB.Add(a, b)
	mulCD.Mul(c, d)
	subABCD.Sub(sumAB, mulCD)
	divE.Div(subABCD, e)
	fh := Hash64(fhp, divE.Bytes())
	finalBigNumber := ZeroBigInt()
	finalBigNumber.SetBytes(fh[:])
	return finalBigNumber
}

// CompletePass represents all the pass in serial plus the final hash arithmetic function
func CompletePass(i []int, data, lastblockhash []byte) (*big.Int, error) {
	passA, errA := HashPassA(data, i[0])
	if errA != nil {
		fmt.Println(errA)
		return nil, errA
	}

	passB, errB := HashPassB(data, i[1])
	if errB != nil {
		fmt.Println(errB)
		return nil, errB
	}

	passC, errC := HashPassC(data, i[2])
	if errC != nil {
		fmt.Println(errC)
		return nil, errC
	}

	passD, errD := HashPassD(data, i[3])
	if errD != nil {
		fmt.Println(errD)
		return nil, errD
	}

	passE, errE := HashPassE(data, i[4])
	if errE != nil {
		fmt.Println(errE)
		return nil, errE
	}
	finalHash := CalcFinalHash(passA, passB, passC, passD, passE)
	return finalHash, nil
}
