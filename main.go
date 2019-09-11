package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"
)

var (
	maxNonce = math.MaxInt64
)

var indexesList [][]int

func main() {

	bc := NewBlockchain()
	i := []int{1, 2, 3, 4, 5}
	indexesList = append(indexesList, i)
	i = GetFuncIndexes(bc.blocks[len(bc.blocks)-1].PrevBlockHash)
	indexesList = append(indexesList, i)
	bc.AddBlock("Send 1 Promethium to HexDev", i)
	i = GetFuncIndexes(bc.blocks[len(bc.blocks)-1].PrevBlockHash)
	indexesList = append(indexesList, i)
	bc.AddBlock("Send 2 Promethium to DSV", i)

	for i, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate(block.PrevBlockHash, indexesList[i])))
		fmt.Println()
	}
}

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte, indexes []int) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run(prevBlockHash, indexes)
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}, []int{1, 2, 3, 4, 5})
}

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string, indexes []int) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, indexes)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(512-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run(prevBlockHash []byte, indexes []int) (int, []byte) {
	var hashInt big.Int
	var hash [64]byte

	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hi, err := CompletePass(indexes, data, prevBlockHash)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		hash, _ = BytesTo64Bytes(hi.Bytes())
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate(prevBlockHash []byte, indexes []int) bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hi, err := CompletePass(indexes, data, prevBlockHash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hashInt.SetBytes(hi.Bytes())
	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
