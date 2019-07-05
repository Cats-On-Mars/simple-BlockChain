package main

import (
	"fmt"
)

func main() {
	//block := newBlock("老师转给班长一枚比特币",[]byte{})
	//fmt.Printf("前区块哈希%x\n",block.PrevHash)
	//fmt.Printf("当前区块哈希%x\n",block.Hash)
	//fmt.Printf("区块数据%s\n",block.Data)


	blockChain := newBlockChain()

	blockChain.addBlock("班长向班花转了3枚比特币")
	blockChain.addBlock("班长向班花转了10枚比特币")

	for i,block := range blockChain.Blocks{
		fmt.Printf("======== 当前区块高度 %d ========\n",i)
		fmt.Printf("前区块哈希:%x\n",block.PrevHash)
		fmt.Printf("当前区块哈希:%x\n",block.Hash)
		fmt.Printf("区块数据:%s\n",block.Data)
	}



}
