package main

import (
	"math/big"
	"testing"
)

func TestIntToHex(t *testing.T) {
	b := IntToHex(1222121)
	t.Log(b)
}

func TestReverseString(t *testing.T) {
	s := "string to reverse"
	t.Log("original string : ", s)
	rs := ReverseString(s)
	t.Log("string reversed : ", rs)
}

func TestReverseBytes(t *testing.T) {
	b := []byte{12, 23, 33, 43}
	t.Log("this is the byte slice to reverse : ", b)
	rb := ReverseBytes(b)
	t.Log("this is the reversed byte slice : ", rb)
}

func TestBytesTo64Bytes(t *testing.T) {
	b := []byte{11, 21, 31, 41, 51, 61, 71, 81}
	t.Log("this is the original byte slice : ", b)
	b64, err := BytesTo64Bytes(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the 64 byte slice : ", b64)
}

func TestRemoveDecFromFloat(t *testing.T) {
	fn := ZeroBigFloat().SetFloat64(1.32344353466644364345)
	f, err := RemoveDecFromFloat(fn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the bigInt with removed dec : ", f)
}

func TestReverseBytes32To64(t *testing.T) {
	b32 := [32]byte{1, 2, 3, 4, 5}
	t.Log("this is the 32 byte slice : ", b32)
	b64Reverse := ReverseBytes32To64(b32)
	t.Log("this is the reversed 64 byte slice : ", b64Reverse)
}

func TestSelectLastDigits(t *testing.T) {
	a := big.NewInt(12345678910)
	t.Log("original big int : ", a)
	i, err := SelectLastDigits(a, 7)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("this is the big int after the last digit selection : ", i)
}
