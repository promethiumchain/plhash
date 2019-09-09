package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
	"unicode/utf8"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return []byte{}
	}

	return buff.Bytes()
}

// ByteToBigInt takes a byte slice as input and returns a big.Int
func ByteToBigInt(in []byte) *big.Int {
	out := new(big.Int)
	out.SetBytes(in)
	return out
}

// BigIntToBigFloat takes a big int as input and returns a float64 with precision given
func BigIntToBigFloat(in *big.Int, prec uint) *big.Float {
	f := new(big.Float)
	f.SetInt(in)
	f.SetPrec(prec)

	return f
}

// ReverseString returns the reverse of the string input
func ReverseString(input string) string {
	size := len(input)
	buf := make([]byte, size)
	for i := 0; i < size; {
		r, n := utf8.DecodeRuneInString(input[i:])
		i += n
		utf8.EncodeRune(buf[size-i:], r)
	}
	return string(buf)
}

// ReverseBytes returns the reversed byte slice
func ReverseBytes(in []byte) []byte {
	revIn := in
	for i := len(revIn)/2 - 1; i >= 0; i-- {
		opp := len(revIn) - 1 - i
		revIn[i], revIn[opp] = revIn[opp], revIn[i]
	}
	return revIn
}

// BytesTo64Bytes returns a filled 64 byte slice from the input
func BytesTo64Bytes(in []byte) ([64]byte, error) {
	var b [64]byte
	i := len(in)
	if i == 0 {
		return [64]byte{}, fmt.Errorf("%s", "size of input is zero")
	}
	for index := 0; index < i; index++ {
		b[index] = in[index]
	}
	return b, nil
}

// RemoveDecFromFloat returns a big.Int from a input float after the deciminal removal
func RemoveDecFromFloat(in *big.Float) (*big.Int, error) {
	s := fmt.Sprintf("%.128f", in)
	if len(s) == 0 {
		return nil, fmt.Errorf("%s", "string has 0 zero length")
	}
	x := strings.Replace(s, ".", "", -1)
	a := big.NewInt(0)
	a.SetString(x, 0)
	return a, nil
}

// ReverseBytes32To64 returns a 64 byte slice filled with a 32 byte and its reverse version
func ReverseBytes32To64(in [32]byte) [64]byte {
	var b [64]byte
	revIn := in
	for i := len(revIn)/2 - 1; i >= 0; i-- {
		opp := len(revIn) - 1 - i
		revIn[i], revIn[opp] = revIn[opp], revIn[i]
	}
	for i := 0; i < 32; i++ {
		b[i] = in[i]
		b[i+32] = revIn[i]
	}
	return b
}

// SelectLastDigits select the number of digits from a big number and returns the outcome
// depth is always a multiply of diff
func SelectLastDigits(in *big.Int, depth uint64) (*big.Int, error) {
	a := ZeroBigInt()
	s := in.String()
	if s == "0" {
		return nil, fmt.Errorf("%s", "string has value of 0? something is not ok")
	}
	fs := strings.TrimRight(s, "0")
	index := int(depth)
	fx := fs[len(fs)-index:]
	a.SetString(fx, 10)
	return a, nil
}

// PrintMessage prints a message between spaces and artifacts
func PrintMessage(args ...interface{}) {
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(args...)
	fmt.Println("---------------------------------------------------------------")
	fmt.Println()
}

// PrintFunctions prints the functions selected by name
func PrintFunctions(indexes []int) {
	for i := range indexes {
		PrintMessage("function selected for pass : ", GetFunctionName(mathFuncs.FuncList[i]))
	}
}
