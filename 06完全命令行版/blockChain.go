package main

import (
	"bolt"
	"fmt"
	"os"
)

const dbfile = "blockChainDb.db"
const blockbucket = "block5"
const lasthash = "LastHash"
const createMsg = "这是山寨版比特币"

//定义区块链结构
type BlockChain struct{
	//Blocks []*Block
	db *bolt.DB
	lastHash []byte
}


//区块链生成函数
func newBlockChain(address string) *BlockChain{
	if IsBlockChainExist(){
		fmt.Println("blockChain already exist!")
		os.Exit(1)
	}
	db, e := bolt.Open(dbfile, 0600, nil)
	checkErr(e,"bolt.Open1")

	var lastHash []byte
	e = db.Update(func(tx *bolt.Tx) error {
		//创建bucket
		bucket, e := tx.CreateBucket([]byte(blockbucket))
		checkErr(e,"tx.CreateBucket")

		//写数据
		coinbase := NewCoinbaseTX(address,createMsg)
		genesis := newGensisBlock(coinbase)

		e = bucket.Put(genesis.Hash, genesis.Serialize())
		checkErr(e,"bucket.Put1")
		e = bucket.Put([]byte(lasthash), genesis.Hash)
		checkErr(e,"bucket.Put2")

		lastHash = genesis.Hash

		return nil
	})
	checkErr(e,"newBlockChain")

	return &BlockChain{db,lastHash}
}

func GetBlockChainHandler()*BlockChain{
	if !IsBlockChainExist(){
		fmt.Println("blockChain already exist!")
		os.Exit(1)
	}
	db, e := bolt.Open(dbfile, 0600, nil)
	checkErr(e,"bolt.Open2")

	var lastHash []byte
	e = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockbucket))
		if bucket != nil{
			lastHash = bucket.Get([]byte(lasthash))
		}
		return nil
	})
	checkErr(e,"GetBlockChainHandler")

	return &BlockChain{db,lastHash}
}

//添加区块
func (bc *BlockChain)addBlock(transaction []*transaction){
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
	block := newBlock(transaction,prevBlockHash)

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

//TODO:？？？
func IsBlockChainExist() bool{
	_, e := os.Stat(dbfile)
	if os.IsNotExist(e){
		return false
	}
	return true
}



func (bc *BlockChain)FindUnspendTransactions(address string) []transaction{
	var transactions []transaction
	var spentUTXOs =make(map[string][]int64) //[交易id][]索引
	bci := bc.Iterator()
	for{
		block := bci.PrevBlock()
		for _,tx := range block.transactions{
			txid := string(tx.TXID)


			//遍历当前交易的输入，找到当前地址已花费的utxos
			for _,input :=range tx.TXInputs{
				if tx.IsCoinbase() == false{
					if input.CaUnlockedByAddress(address){
						spentUTXOs[txid]=append(spentUTXOs[txid],input.ReferOutputIndex)
					}
				}
			}

		LABEL1:
			//遍历当前交易的输出，找到满足条件的输出
			for outputIndex,output := range tx.TXOutputs{
				if spentUTXOs[txid] != nil{
					for _,usedIndex := range spentUTXOs[txid]{
						if int64(outputIndex) == usedIndex{
							continue LABEL1 //TODO:跳出到label1
						}
					}
				}
				if output.CaBeUnlockedByAddress(address){
					transactions = append(transactions,*tx)
				}
			}

		}

		if len(block.PrevHash)==0{
			break
		}
	}
	return transactions
}


func (bc *BlockChain)FindUTXOs(address string)[]Output{
	var outputs []Output
	txs := bc.FindUnspendTransactions(address)
	for _,tx := range txs{
		for _,output := range tx.TXOutputs{
			outputs = append(outputs,output)
		}
	}
	return outputs
}


func (bc *BlockChain)FindSuitableUTXOs(address string,amount float64) (float64,map[string][]int64){
	txs := bc.FindUnspendTransactions(address)
	var countTotal float64
	var container =make(map[string][]int64)
LABEL2:
	for _,tx := range txs{
		for index,output := range tx.TXOutputs{
			if countTotal < amount {
				if output.CaBeUnlockedByAddress(address) {
					countTotal += output.Value
					container[string(tx.TXID)] = append(container[string(tx.TXID)],int64(index))
				}
			}else{
				break LABEL2
			}
		}
	}
	return countTotal,container
}
