package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(Data string) *Block {
	newBlock := Block{Data, "", getLastHash()}
	Hash := sha256.Sum256([]byte(newBlock.Data + newBlock.PrevHash))
	newBlock.Hash = fmt.Sprintf("%x", Hash)
	return &newBlock
}

func (b *blockchain) AddBlock(Data string) {
	b.blocks = append(b.blocks, createBlock(Data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Data")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
