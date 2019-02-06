package core

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

// Block defines the structure of a single block
type Block struct {
	Timestamp     int64   // Timestamp that the block gets created
	Data          []byte  // Actual valuable information containing in the block
	PrevBlockHash []byte  // Hash of previous block
	Hash          []byte  // Hash of current block
}

// SetHash is the function to set the value of the block hash based on 
// timestamp, prevBlockHash, and data
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock is the function to create a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}