package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
	addBlock --data DATA	"add a block to blockchain"
	printChain	"print all blocks"
`

const AddBlockCmdString = "addBlock"
const PrintChainCmdString = "printChain"

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("Invalid input!")
	fmt.Println(usage)
	os.Exit(1)
}

func (cli *CLI) parameterCheck() {
	if len(os.Args) < 2 {
		cli.printUsage()
	}
}

func (cli *CLI) Run() {
	cli.parameterCheck()

	addBlockCmd := flag.NewFlagSet(AddBlockCmdString, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString, flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data", "", "block transaction info!")

	switch os.Args[1] {
	case AddBlockCmdString:
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("Run1()", err)
		if addBlockCmd.Parsed() {
			if *addBlockCmdPara == "" {
				fmt.Println("addBlock data should not be empty!")
				cli.printUsage()
			}

			cli.AddBlock(*addBlockCmdPara)
		}

	case PrintChainCmdString:
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr("Run2()", err)
		if printChainCmd.Parsed() {
			cli.PrintChain()
		}

	default:
		cli.printUsage()
	}
}
