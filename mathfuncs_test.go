package main

import (
	"testing"
)

func TestGetFuncIndexes(t *testing.T) {
	lastBlockHash := "0x43b60820064b41e64491136cc333ca5862e855eb03bdbcd67a6f70ce4478b100"
	t.Log("this is the last block hash : ", lastBlockHash)
	indexes := GetFuncIndexes([]byte(lastBlockHash))
	t.Log("next block math functions where selected with the latest block hash, numbers are : ", indexes)

}

func TestNewFuncList(t *testing.T) {
	nfl := NewFuncList()
	for i, f := range nfl.FuncList {
		name := GetFunctionName(f)
		t.Log("Func index : ", i, " and func name : ", name)
	}
}
