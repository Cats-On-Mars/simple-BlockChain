package main

import (
	"fmt"
	"os"
)

func (cli *CLI)CreateChain(address string){
	bc := newBlockChain(address)
	defer bc.db.Close()
	fmt.Println("create blockChain successfully!!")
}


//func (cli *CLI)AddBlock(data string){
//	cli.bc.addBlock(data)
//	fmt.Println("Add Block Succeed!")
//}

func (cli *CLI)PrintChain(){
	bc := GetBlockChainHandler()
	iterator := bc.Iterator()
	if len(iterator.currentHash) == 0{
		fmt.Println("len(iterator.currentHash) == 0")
		os.Exit(1)
	}

	for{
		block := iterator.PrevBlock()
		//fmt.Printf("======== 当前区块高度 %d ========\n",i)
		fmt.Printf("Version:%d\n",block.Version)
		fmt.Printf("PreBlockHash:%x\n",block.PrevHash)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("TimeStamp:%d\n",block.TimeStamp)
		fmt.Printf("MerkleRoot:%x\n",block.MerkleRoot)
		fmt.Printf("TargetBits:%d\n",block.TargetBits)
		fmt.Printf("Nonce:%d\n",block.Nonce)

		fmt.Printf("transactions:%v\n",block.transactions)

		pow := NewProofOfWork(block)
		fmt.Printf("IsValid:%v\n",pow.IsValid())
		if len(iterator.currentHash) == 0{
			//fmt.Println("len(iterator.currentHash) == 0, Print finshed")
			os.Exit(1)
		}
	}

}

func (cli *CLI)GetBalance(address string){
	bc := GetBlockChainHandler()
	defer bc.db.Close()

	var balance float64
	outputs := bc.FindUTXOs(address)
	for _,output := range outputs{
		balance += output.Value
	}
	fmt.Printf("%s has %v BTC",address,balance)
}