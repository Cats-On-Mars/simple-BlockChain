package main

//定义区块链结构
type BlockChain struct{
	Blocks []*Block
}

//创世区块生成函数
func gensisBlock() *Block{
	return newBlock("这是创世区块",[]byte{})
}

//区块链生成函数
func newBlockChain() *BlockChain{
	genesisblock := gensisBlock()
	return &BlockChain{
		Blocks:[]*Block{genesisblock},
	}

}


//添加区块
func (bc *BlockChain)addBlock(data string){
	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	block := newBlock(data, lastBlock.Hash)

	bc.Blocks = append(bc.Blocks, block)

}