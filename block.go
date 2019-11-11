package main

import (
	"bytes"
	"encoding/gob"
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
	Transaction   []*Transaction
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	CheckErr("Serialize", err)
	return buffer.Bytes()

}
func Deserialize(data []byte) *Block {
	if len(data) == 0 {
		return nil
	}
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	CheckErr("Deserialize", err)
	return &block
}
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		MerKelRoot:    []byte{},
		TimeStamp:     time.Now().Unix(),
		Bits:          targetBits,
		Transaction:   txs}

	pow := NewProofOfwork(&block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}
