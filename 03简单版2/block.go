package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//定义区块结构
type Block struct{
	Version uint64
	PrevHash []byte
	TimeStamp uint64
	MerkleRoot []byte
	TargetBits uint64
	Nonce uint64

	Hash []byte   //为了方便实现做的简化，正常比特币区块不包含自己的哈希值
	Data []byte  //模拟区块体，交易
}

//区块生成工厂函数
func newBlock(data string,prevBlockHash []byte) *Block{
	block := Block{
		Version:00,
		PrevHash:prevBlockHash,
		TimeStamp:uint64(time.Now().Unix()),
		MerkleRoot:[]byte{},
		TargetBits:0,
		Nonce:0,

		Hash:[]byte{},
		Data:[]byte(data),
	}
	block.setHash()
	return &block
}

//哈希生成函数
func (block *Block)setHash(){
	//拼装数据方法1：append方法
	var blockInfo []byte
	blockInfo = append(blockInfo,[]byte(string(block.Version))...)
	blockInfo = append(blockInfo,block.PrevHash...)
	blockInfo = append(blockInfo,[]byte(string(block.TimeStamp))...)
	blockInfo = append(blockInfo,block.MerkleRoot...)
	blockInfo = append(blockInfo,[]byte(string(block.TargetBits))...)
	blockInfo = append(blockInfo,[]byte(string(block.Nonce))...)
	blockInfo = append(blockInfo,block.Data...)

	//拼装数据方法2：join方法
	temp := [][]byte{
		IntToByte(block.Version),
		block.PrevHash,
		IntToByte(block.TimeStamp),
		block.MerkleRoot,
		IntToByte(block.TargetBits),
		IntToByte(block.Nonce),
		block.Data,
	}
	blockInfo = bytes.Join(temp, []byte{})

	//sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//创世区块生成函数
func newGensisBlock() *Block{
	return newBlock("这是创世区块",[]byte{})
}
