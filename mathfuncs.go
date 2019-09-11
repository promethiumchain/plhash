package main

import (
	"math/big"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// MathFuncs represent the collection of the math functions used by the chain hasher
type MathFuncs struct {
	FuncList []func(*big.Float) *big.Float
}

// NewFuncList returns a new list of math functions
func NewFuncList() *MathFuncs {
	return &MathFuncs{
		FuncList: createMathList(),
	}
}

func createMathList() []func(*big.Float) *big.Float {
	var collection []func(*big.Float) *big.Float
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)

	//Dummy Append
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)

	return collection
}

// GetFuncIndexes retuns the indexes for the math funcs chosen
func GetFuncIndexes(lastBlockHash []byte) []int {
	var funcIndexes = make([]int, 5)
	a := big.NewInt(0)
	a.SetBytes(lastBlockHash)
	// Trim ending zeros
	s := a.String()
	if s == "0" {
		return []int{1, 2, 3, 4, 5}
	}
	st := strings.TrimRight(s, "0")
	sm := st[len(st)-10:]
	// Seperate values and store them
	sm1, _ := strconv.Atoi(sm[:2])
	sm2, _ := strconv.Atoi(sm[2:4])
	sm3, _ := strconv.Atoi(sm[4:6])
	sm4, _ := strconv.Atoi(sm[6:8])
	sm5, _ := strconv.Atoi(sm[8:10])
	funcIndexes[0] = sm1
	funcIndexes[1] = sm2
	funcIndexes[2] = sm3
	funcIndexes[3] = sm4
	funcIndexes[4] = sm5
	PrintFunctions(funcIndexes)
	return funcIndexes
}

// GetFunctionName returns the name of the function
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
