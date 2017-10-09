package main

//  Blockchain: ordered, back-linked list. Blocks are stored
//  in the insertion order such that each block is linked to the
//  previous one
//  Implementation using an Array and a Map.
//  Array - will maintain ordered hashes
//  Map - would keep unorderd hash → block pairs

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// A new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
