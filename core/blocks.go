package core

import (
	"bytes"
	"encoding/gob"
	"log"
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

// Serialize converts Block into byte array
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

// Deserialize converts byte array back to Block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}