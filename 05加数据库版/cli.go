package main

import (
	"flag"
	"fmt"
	"os"
)

const Usage =`
	./block addBlock --data DATA   "add a block to blockChain"
	./block printChain             "print all blocks"
`

type CLI struct{
	bc *BlockChain
}

func (cli *CLI)Run()  {
	if len(os.Args) <2{
		fmt.Println("too few parameters!",Usage)
		os.Exit(1)
	}

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data","","block info")

	switch os.Args[1]{
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		checkErr(err,"addBlockCmd.Parse")
		if addBlockCmd.Parsed(){
			if *addBlockCmdPara == ""{
				fmt.Println("data is empty")
				os.Exit(1)
			}else{
				//cli.bc.addBlock(*addBlockCmdPara)
				cli.AddBlock(*addBlockCmdPara)
			}
		}
	case "printChain":
		//fmt.Println("开始遍历")
		err := printCmd.Parse(os.Args[2:])
		checkErr(err,"addBlockCmd.Parse")
		//fmt.Println("命令行已解析")
		if printCmd.Parsed(){
			//fmt.Println("命令行已验证解析")
			cli.PrintChain()
		}
	default:
		fmt.Println("invalid cmd\n",Usage)
		os.Exit(1)
	}
}


//func (cli *CLI)Run(){
//	var blockInfo string
//	&blockInfo = flag.String("addblock", "", "addBlock")
//	flag.Parse()
//	if blockInfo == ""{
//
//	}
//
//}