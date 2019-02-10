package core

import (
	"log"

	"github.com/boltdb/bolt"
)

const (
	dbFile = "blockchain_%s.db"
	blocksBucket = "blocks"
)

// Blockchain defines the structure of a block-chain
type Blockchain struct {
	tip []byte     // tip of block hash stored in DB
	db  *bolt.DB   // DB connection
}

// AddBlock add a new block to a blockchain
func (bc *Blockchain) AddBlock(data string) {
	var prevHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		prevHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, prevHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// NewGenesisBlock creates the first block for a blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{tip, db}

	return &blockchain
}