package main

import (
	"fmt"
)

func main() {
	//block := newBlock("老师转给班长一枚比特币",[]byte{})
	//fmt.Printf("前区块哈希%x\n",block.PrevHash)
	//fmt.Printf("当前区块哈希%x\n",block.Hash)
	//fmt.Printf("区块数据%s\n",block.Data)


	bc := newBlockChain()

	bc.addBlock("班长向班花转了3枚比特币")
	bc.addBlock("班长向班花转了10枚比特币")

	for i,block := range bc.Blocks{
		fmt.Printf("======== 当前区块高度 %d ========\n",i)
		fmt.Printf("Version:%d\n",block.Version)
		fmt.Printf("PreBlockHash:%x\n",block.PrevHash)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("TimeStamp:%d\n",block.TimeStamp)
		fmt.Printf("MerkleRoot:%x\n",block.MerkleRoot)
		fmt.Printf("TargetBits:%d\n",block.TargetBits)
		fmt.Printf("Nonce:%d\n",block.Nonce)

		fmt.Printf("区块数据:%s\n",block.Data)


	}



}
