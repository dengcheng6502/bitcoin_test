package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
}

func (cli *CLI) PrintChain() {
	it := cli.bc.NewIterator()

	for {

		block := it.Next()
		fmt.Printf("Version: %d\n", block.Version)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("Bits: %d\n", block.Bits)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Isvalid: %v\n", NewProofOfwork(block).IsValid())
		if len(block.PrevBlockHash) == 0 {
			fmt.Println("print over!")
			break
		}

	}
}
