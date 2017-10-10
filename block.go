//  Blockchain, a decentralised, timestamped public ledger.
//  Block, conatins technical information like
//    - Timestamp (current timestamp when block is created),
//    - Data (Informations like version)
//    - Hash (Hash of its own block)
//    - Hash of previous block (Hash of prev block)
//  as a block header and data.

package main

import (
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Hash calcualtion - it should be computationally difficult operation
// to compute a hash to makes hard to add a new block into blockchian,
// thus preventing modification of block

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis block: To add a new block we need an existing block,
// So, in any blockchain, there must be at least one block, such block
// is called Genesis block

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
