//  Blockchain, a decentralised, timestamped public ledger.
//  Block, conatins technical information like
//    - Timestamp (current timestamp when block is created),
//    - Data (Informations like version)
//    - Hash (Hash of its own block)
//    - Hash of previous block (Hash of prev block)
//  as a block header and data.

package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Hash calcualtion - it should be computationally difficult operation
// to compute a hash to makes hard to add a new block into blockchian,
// thus preventing modification of block

// SetHash - Concatenate block fields then calcualte SHA-256

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// Genesis block: To add a new block we need an existing block,
// So, in any blockchain, there must be at least one block, such block
// is called Genesis block

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
