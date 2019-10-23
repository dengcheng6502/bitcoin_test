package src

import (
	"time"
)

type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte
	MerKelRoot    []byte
	TimeStamp     int64
	Bits          int64
	Nonce         int64
	Data          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		MerKelRoot:    []byte{},
		TimeStamp:     time.Now().Unix(),
		Bits:          targetBits,
		Data:          []byte(data)}

	pow := NewProofOfwork(&block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
