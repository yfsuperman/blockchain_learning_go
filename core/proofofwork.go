package core

import (
	"bytes"
	"math"
	"math/big"
	"fmt"
	"crypto/sha256"

	"github.com/yfsuperman/blockchain_learning_go/utils"
)

const (
	targetBits = 24
)

var (
	maxCounter = math.MaxInt64
)

// ProofOfWork defines a structure to help validate block in a chain
// It consistes of a block and a target number help limit the upper
// boundary of the hash
type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepData(counter int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(counter)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	counter := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for counter < maxCounter {
		data := pow.prepData(counter)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			counter++
		}
	}
	fmt.Print("\n\n")

	return counter, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepData(pow.block.Counter)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}