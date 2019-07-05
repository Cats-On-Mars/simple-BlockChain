package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

//定义区块结构
type Block struct{
	Version int64
	PrevHash []byte
	TimeStamp int64
	MerkleRoot []byte
	TargetBits int64
	Nonce int64

	Hash []byte   //为了方便实现做的简化，正常比特币区块不包含自己的哈希值
	Data []byte  //模拟区块体，交易
}

func (block *Block)Serialize()[]byte{
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(block)
	checkErr(err,"encoder.Encode")
	return buf.Bytes()
}

func Deserialize(data []byte)*Block{
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	err := decoder.Decode(&block)
	checkErr(err,"decoder.Decode")
	return &block
}

//区块生成工厂函数
func newBlock(data string,prevBlockHash []byte) *Block{
	block := Block{
		Version:00,
		PrevHash:prevBlockHash,
		TimeStamp:time.Now().Unix(),
		MerkleRoot:[]byte{},
		TargetBits:targetBits,
		Nonce:0,

		Hash:[]byte{},
		Data:[]byte(data),
	}
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	return &block
}



//创世区块生成函数
func newGensisBlock() *Block{
	return newBlock("这是创世区块",[]byte{})
}
