package main

import (
	"bytes"
	"crypto/sha256"
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
		Bits:          1,
		Nonce:         1,
		Data:          []byte(data)}

	block.SetHash()
	return &block
}

func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}

	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
