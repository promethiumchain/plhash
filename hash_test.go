package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	Init()
	// indexes, _ := GetFuncIndexes([]byte("0x43b60820064b41e64491136cc333ca5862e855eb03bdbcd67a6f70ce4478b100"))
	data := []byte("This is the genesis example")
	// h32 := sha3.New256()
	// out32 := Hash32(h32, data)
	// t.Log("this is the 32 byte slice : ", out32)

	// h64 := sha3.New512()
	// out64 := Hash64(h64, data)
	// t.Log("this is 64 byte slice : ", out64)

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

}

func TestMathFuncs(t *testing.T) {
	Init()
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
	Init()

	bc := NewBlockchain()
	bc.AddBlock("Send 1 Promethium to HexDev")
	bc.AddBlock("Send 2 Promethium to DSV")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

func TestDiffString(t *testing.T) {
	difflevel := 10
	s := ConstractDiffString(difflevel)
	t.Log(s)
}
