package main

import "os"

//定义区块链结构
type BlockChain struct{
	Blocks []*Block
}


//区块链生成函数
func newBlockChain() *BlockChain{
	genesisblock := newGensisBlock()
	return &BlockChain{
		Blocks:[]*Block{genesisblock},
	}

}

//添加区块
func (bc *BlockChain)addBlock(data string){
	if len(bc.Blocks) <= 0{        //避免出现越界
		os.Exit(1)
	}

	lastBlock := bc.Blocks[len(bc.Blocks)-1]   //取出最后一个区块，得到PrevBlockHash

	block := newBlock(data, lastBlock.Hash)

	bc.Blocks = append(bc.Blocks, block)

}