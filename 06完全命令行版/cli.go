package main

import (
	"flag"
	"fmt"
	"os"
)

const Usage =`
	createChain -address Address   "create block chain"
	addBlock -data DATA   "add a block to blockChain"
	printChain             "print all blocks"
	getBalance -address Address "get the balance"
`

type CLI struct{
	//bc *BlockChain
}

func (cli *CLI)Run()  {
	if len(os.Args) <2{
		fmt.Println("too few parameters!",Usage)
		os.Exit(1)
	}

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	createChainCmd := flag.NewFlagSet("createChain",flag.ExitOnError)
	printCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data","","block info")
	createChainCmdPara := createChainCmd.String("address","","create address")
	getBalanceCmdPara := getBalanceCmd.String("address","","user address")

	switch os.Args[1]{
	case "createChain":
		err := createChainCmd.Parse(os.Args[2:])
		checkErr(err,"createChainCmd.Parse")
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		checkErr(err,"addBlockCmd.Parse")
	case "printChain":
		err := printCmd.Parse(os.Args[2:])
		checkErr(err,"addBlockCmd.Parse")
	case "getBalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		checkErr(err,"getBalanceCmd.Parsed")
	default:
		fmt.Println("invalid cmd\n",Usage)
		os.Exit(1)
	}

	if createChainCmd.Parsed(){
		if *createChainCmdPara == ""{
			fmt.Println("address is empty")
			os.Exit(1)
		}
		cli.CreateChain(*createChainCmdPara)
	}

	if addBlockCmd.Parsed(){
		if *addBlockCmdPara == ""{
			fmt.Println("data is empty")
			os.Exit(1)
		}
		//cli.AddBlock(*addBlockCmdPara)
	}

	if printCmd.Parsed(){
		cli.PrintChain()
	}

	if getBalanceCmd.Parsed(){
		if *getBalanceCmdPara == ""{
			fmt.Println("address is empty")
			os.Exit(1)
		}
		cli.GetBalance(*getBalanceCmdPara)
	}
}


