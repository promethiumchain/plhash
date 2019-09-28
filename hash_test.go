package main

import (
	"crypto/sha512"
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"golang.org/x/crypto/sha3"
)

func TestHash(t *testing.T) {
	data := []byte("This is the genesis example")

	passA, err := HashPassA(data, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the outcome via pass A -> SHA256 : ", passA)

	passB, err := HashPassB(data, 7)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the outcome via pass B -> SHA3.512 : ", passB)

	passC, err := HashPassC(data, 12)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the outcome via pass C -> SHA3.512 : ", passC)

	passD, err := HashPassD(data, 8)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the outcome via pass D -> BLAKE2S : ", passD)

	passE, err := HashPassE(data, 15)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the outcome via pass E -> BLAKE2B : ", passE)

	finalHash := CalcFinalHash(passA, passB, passC, passD, passE)
	t.Log("final hash big number : ", finalHash)
	t.Log("final hash hex : ", toHexInt(finalHash))
}

func TestSha(t *testing.T) {
	h1 := sha3.New512()
	b1 := []byte("This is the genesis example")
	fh1 := Hash64(h1, b1)
	finalBigNumber1 := ZeroBigInt()
	finalBigNumber1.SetBytes(fh1[:])
	h2 := sha3.New512()
	b2 := []byte("This is the genesis example")
	fh2 := Hash64(h2, b2)
	h3 := sha512.New()
	fh3 := Hash64(h3, b2)
	finalBigNumber3 := ZeroBigInt()
	finalBigNumber3.SetBytes(fh3[:])
	finalBigNumber2 := ZeroBigInt()
	finalBigNumber2.SetBytes(fh2[:])
	t.Log(toHexInt(finalBigNumber1))
	t.Log(toHexInt(finalBigNumber2))
	t.Log(finalBigNumber1)
	t.Log(finalBigNumber2)
	t.Log("this is the double pass : ", finalBigNumber3)

	if fh1 == fh3 {
		t.Log("data are the same")
	} else {
		t.Log("data dont match")
	}
}

func toHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n)
}

func TestMathFuncs(t *testing.T) {
	data := []byte("This is the genesis example")

	passA, errA := HashPassA(data, 1)
	if errA != nil {
		t.Fatal(errA)
	}
	t.Log("this is the outcome via pass A -> SHA256 : ", passA)
	passB, errB := HashPassB(data, 2)
	if errB != nil {
		t.Fatal(errB)
	}
	t.Log("this is the outcome via pass B -> SHA3.512 : ", passB)
	passC, errC := HashPassC(data, 3)
	if errC != nil {
		t.Fatal(errC)
	}
	t.Log("this is the outcome via pass C -> SHA3.512 : ", passC)
	passD, errD := HashPassD(data, 4)
	if errD != nil {
		t.Fatal(errA)
	}
	t.Log("this is the outcome via pass D -> BLAKE2S : ", passD)
	passE, errE := HashPassE(data, 5)
	if errE != nil {
		t.Fatal(errE)
	}
	t.Log("this is the outcome via pass E -> BLAKE2B : ", passE)

}

func TestPow(t *testing.T) {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 Promethium to HexDev", []int{1, 2, 3, 4, 5})
	bc.AddBlock("Send 2 Promethium to DSV", []int{1, 2, 3, 4, 5})

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate(block.PrevBlockHash, []int{1, 2, 3, 4, 5})))
		fmt.Println()
	}
}
