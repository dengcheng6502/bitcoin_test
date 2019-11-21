package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"os"
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
	if len(tx.TXInputs) == 1 {
		if len(tx.TXInputs[0].TXID) == 0 && tx.TXInputs[0].Vout == -1 {
			return true
		}
	}
	return false
}

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {

	validUTXOs := make(map[string][]int64)
	var total float64
	validUTXOs, total = bc.FindSuitableUTXOs(from, amount)

	if total < amount {
		fmt.Println("Not enough money!")
		os.Exit(1)
	}

	var inputs []TXInput
	var outputs []TXOutput

	for txId, outputIndexes := range validUTXOs {
		for _, index := range outputIndexes {
			input := TXInput{[]byte(txId), int64(index), from}
			inputs = append(inputs, input)

		}
	}

	output := TXOutput{amount, to}
	outputs = append(outputs, output)

	if total > amount {
		output := TXOutput{total - amount, to}
		outputs = append(outputs, output)
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetTXID()
	return &tx

}
