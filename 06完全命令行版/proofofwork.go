package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//tODO:学习big包

const targetBits = 24  //设置目标难度值的位数？？？

type ProofOfWork struct{
	Block *Block
	TargetBit *big.Int
}

//构造一个工作量证明，主要就是计算出难度
func NewProofOfWork(block *Block)*ProofOfWork{
	var IntTarget = big.NewInt(1)
	IntTarget.Lsh(IntTarget,uint(256-targetBits))
	return &ProofOfWork{block,IntTarget}
}

func (pow *ProofOfWork)prepareRawData(nonce int64)[]byte{
	block := pow.Block
	temp := [][]byte{
		IntToByte(block.Version),
		block.PrevHash,
		IntToByte(block.TimeStamp),
		block.MerkleRoot,
		IntToByte(targetBits),
		IntToByte(nonce),
		//block.Data,
		//block.transactions,  //TODO:
	}
	data := bytes.Join(temp, []byte{})
	return data
}

func (pow *ProofOfWork)Run()(int64,[]byte){
	var nonce int64
	var hash [32]byte
	var hashInt big.Int

	fmt.Println("Begin Mining...")
	fmt.Printf("target hash:   %x\n",pow.TargetBit.Bytes())
	for nonce< math.MaxInt64{
		data := pow.prepareRawData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.TargetBit)==-1{
			fmt.Printf("found Hash:%x\n",hash)
			break
		}else{
			nonce++
		}

		//intHash := ByteToInt(hash[:])
		//bigIntHash := big.NewInt(intHash)
	}
	return nonce,hash[:]
}

func (pow *ProofOfWork)IsValid()bool{
	data := pow.prepareRawData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	var IntHash big.Int
	IntHash.SetBytes(hash[:])
	return IntHash.Cmp(pow.TargetBit)==-1

}