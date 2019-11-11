package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

const reward = 12.5

type Transaction struct {
	TXID      []byte
	TXInputs  []TXInput
	TXOutputs []TXOutput
}

type TXInput struct {
	TXID      []byte
	Vout      int64
	ScriptSig string
}

func (input *TXInput) CanUnlockUTXOWith(unlockData string) bool {
	return input.ScriptSig == unlockData
}

type TXOutput struct {
	Value        float64
	ScriptPubKey string
}

func (output *TXOutput) CanBeUnlockWith(unlockData string) bool {
	return output.ScriptPubKey == unlockData

}

func (tx *Transaction) SetTXID() {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)

	CheckErr("SetTXID", err)
	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]
}

func NewCoinbaseTx(address string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("reward to %s %d btc", address, reward)
	}
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	tx := Transaction{
		[]byte{}, []TXInput{input}, []TXOutput{output},
	}
	tx.SetTXID()
	return &tx
}

func (tx *Transaction) IsCoinbase() bool {

}

func (bc *BlockChain) FindUTXOTransactions(address string) []Transaction {
	var UTXOTransactions []Transaction
	it := bc.NewIterator()
	for {
		block := it.Next()
		for _, tx := range block.Transaction {
			for _, output := range tx.TXOutputs {
				if output.CanBeUnlockWith(address) {
					UTXOTransactions = append(UTXOTransactions, *tx)
				}
			}
		}
	}
}

/*
func NewTransaction(from,to string,amount float64.bc *BlockChain){
}
*/
