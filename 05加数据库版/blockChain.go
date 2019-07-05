package main

import (
	"bolt"
	"fmt"
	"os"
)

const dbfile = "blockChainDb.db"
const blockbucket = "block"
const lasthash = "LastHash"

//定义区块链结构
type BlockChain struct{
	//Blocks []*Block
	db *bolt.DB
	lastHash []byte
}


//区块链生成函数
func newBlockChain() *BlockChain{
	db, e := bolt.Open(dbfile, 0600, nil)
	checkErr(e,"bolt.Open")

	var lastHash []byte
	e = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockbucket))
		if bucket != nil{
			//读取lastHash即可
			lastHash = bucket.Get([]byte(lasthash))
		}else{
			//创建bucket
			//写数据
			genesis := newGensisBlock()
			bucket, e := tx.CreateBucket([]byte(blockbucket))
			checkErr(e,"tx.CreateBucket")

			e = bucket.Put(genesis.Hash, genesis.Serialize())
			checkErr(e,"bucket.Put1")
			e = bucket.Put([]byte(lasthash), genesis.Hash)
			checkErr(e,"bucket.Put2")

			lastHash = genesis.Hash
		}

		return nil
	})
	checkErr(e,"db.Update")

	return &BlockChain{db,lastHash}
}

//添加区块
func (bc *BlockChain)addBlock(data string){
	//获取lastHash
	var prevBlockHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockbucket))
		lastHash := bucket.Get([]byte(lasthash))
		prevBlockHash = lastHash
		return nil
	})
	checkErr(err,"bc.db.View")

	//获取到的lastHash作为新区块的prevHash，构造出新区块
	block := newBlock(data,prevBlockHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockbucket))
		err := bucket.Put(block.Hash, block.Serialize())   //将新区块添加到数据库
		checkErr(err,"bucket.Put3")
		err = bucket.Put([]byte(lasthash), block.Hash)     //将新区块的hash作为键lasthash的值
		checkErr(err,"bucket.Put4")
		bc.lastHash = block.Hash
		return nil
	})
	checkErr(err,"bc.db.Update")
}


//构造迭代器结构 区块链+当前哈希
type blockChainIterator struct{
	db *bolt.DB
	currentHash []byte
}

//生成区块链的迭代器，区块链的句柄+区块链的lastHash
func (bc *BlockChain)Iterator()*blockChainIterator{
	return &blockChainIterator{bc.db,bc.lastHash}
}

//（迭代区块链）获取迭代器当前哈希指向的区块
func (it *blockChainIterator)PrevBlock()*Block{
	var block *Block
	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockbucket))
		if bucket == nil{
			fmt.Println("bucket is nil")
			os.Exit(1)
		}

		blockTemp := bucket.Get(it.currentHash)
		block = Deserialize(blockTemp)
		it.currentHash = block.PrevHash
		return nil
	})
	checkErr(err,"prevBlock")
	return block
}