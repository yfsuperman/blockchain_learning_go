package core

import (
	"time"
)

// Block defines the structure of a single block
type Block struct {
	Timestamp     int64   // Timestamp that the block gets created
	Data          []byte  // Actual valuable information containing in the block
	PrevBlockHash []byte  // Hash of previous block
	Hash          []byte  // Hash of current block
	Counter       int     // cryptographic term, used for validate block
}

// NewBlock is the function to create a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	counter, hash := pow.Run()

	block.Hash = hash[:]
	block.Counter = counter

	return block
}