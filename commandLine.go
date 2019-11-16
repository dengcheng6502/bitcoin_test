package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
	createChain --address ADDRESS "create a blockchain"
	send --from FROM --to TO --amount AMOUNT --miner MINER "send coin from FROM to TO"
	addBlock --data DATA	"add a block to blockchain"
	getbalance --address ADRESS "get balance of the address"
	printChain	"print all blocks"
`

const AddBlockCmdString = "addBlock"
const PrintChainCmdString = "printChain"
const CreateChainCmdString = "createChain"
const SednCmdString = "send"
const GetBalanceCmdString = "getbalance"

type CLI struct {
	//bc *BlockChain
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

	CreateChainCmd := flag.NewFlagSet(CreateChainCmdString, flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet(AddBlockCmdString, flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet(GetBalanceCmdString, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString, flag.ExitOnError)

	createChainCmdPara := CreateChainCmd.String("address", "", "address info")
	addBlockCmdPara := addBlockCmd.String("data", "", "block transaction info!")
	getBalanceCmdPara := getBalanceCmd.String("address", "", "address info")
	switch os.Args[1] {
	case CreateChainCmdString:
		err := CreateChainCmd.Parse(os.Args[2:])
		CheckErr("Run0()", err)
		if CreateChainCmd.Parsed() {
			if *createChainCmdPara == "" {
				fmt.Println("address should not be empty")
				cli.printUsage()
			}

			cli.CreateChain(*createChainCmdPara)
		}
	case AddBlockCmdString:
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("Run1()", err)
		if addBlockCmd.Parsed() {
			if *addBlockCmdPara == "" {
				fmt.Println("addBlock data should not be empty!")
				cli.printUsage()
			}

			//cli.AddBlock(*addBlockCmdPara)
		}

	case PrintChainCmdString:
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr("Run2()", err)
		if printChainCmd.Parsed() {
			//cli.PrintChain()
		}

	case GetBalanceCmdString:
		err := getBalanceCmd.Parse(os.Args[2:])
		CheckErr("Run2()", err)
		if getBalanceCmd.Parsed() {
			if *getBalanceCmdPara == "" {
				fmt.Println("address data should not be emptyal!")
				cli.printUsage()
			}

			cli.GetBalance(*getBalanceCmdPara)
		}

	default:
		cli.printUsage()
	}
}
