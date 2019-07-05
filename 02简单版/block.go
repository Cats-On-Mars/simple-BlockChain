package main

import "crypto/sha256"

//定义区块结构
type Block struct{
	PrevHash []byte
	Hash []byte
	Data []byte
}

//区块生成工厂函数
func newBlock(data string,prevBlockHash []byte) *Block{
	block := Block{
		PrevHash:prevBlockHash,
		Hash:[]byte{},
		Data:[]byte(data),
	}
	block.setHash()
	return &block
}

//哈希生成函数
func (block *Block)setHash(){
	//拼装数据
	blockInfo := append(block.PrevHash,block.Data...)
	//sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
