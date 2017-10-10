// Proof of work - Hard work + Prove (do hard work and prove)
// Miner have to do hard work in order to add block into blockchain.
// and get reward for their work after proving.

// Block rate - 6 blocks per hour
// Difficulty increase based on time and miners to maintain consistant block rate

// Hasing - Getting a hash of data
// Hash Fucntion - A function that take data input and return a fix sized Hash.
//    - Original data cannot be restored from a hash. Thus, hashing is not encryption
//    - Hash is unique
//    - Changing a single bit in data will give you complete new Hash
//    - It used to check consistency of block

// POW requirement - first 20 bits of a hash must be zeros
// (Changes based on time and miners)

// Brute force algorithm
// A HASH = block header/ data + counter(increate counter if needed from 0)

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// Difficulty of mining should be adjust based on pow requirement
// less then 256 bit in memory
const targetBits = 18

// Pointer to a block & target
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// Hash(data + counter) -> big integer and check if less then target
// if smaller then target then valid proof
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

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
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
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
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
