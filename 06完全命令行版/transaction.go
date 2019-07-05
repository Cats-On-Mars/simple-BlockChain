package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

const reward float64 =12.5

type transaction struct{
	TXID []byte
	TXInputs []Input
	TXOutputs []Output
}


type Input struct{
	Txid []byte
	ReferOutputIndex int64
	UnlockScript string   //scriptSig
}

type Output struct {
	Value float64
	LockScript string     //scriptPubKey
}

func (tx *transaction)SetTXID(){
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tx)
	checkErr(err,"encoder.Encode")

	hash := sha256.Sum256(buf.Bytes())
	tx.TXID = hash[: ]
}


func (input *Input)CaUnlockedByAddress(unlockdata string)bool{
	return input.UnlockScript == unlockdata
}

func (output *Output)CaBeUnlockedByAddress(unlockdata string)bool{
	return output.LockScript == unlockdata
}

func NewCoinbaseTX(address string,createMsg string) *transaction{
	if createMsg == ""{
		createMsg = fmt.Sprintf("(这是创世区块的input解锁脚本)创世区块创建时间：%v\n", time.Now().Unix())
	}

	input := Input{nil,-1,createMsg}
	output := Output{reward,address}

	tx := transaction{nil,[]Input{input},[]Output{output}}
	tx.SetTXID()
	fmt.Println("CoinbaseTX创建成功")
	return &tx
}

func NewTransaction(from,to string, amount float64,bc *BlockChain) *transaction{
	countTotal,container := bc.FindSuitableUTXOs(from, amount)
	if countTotal < amount{
		fmt.Println("Not Enough Founds")
	}

	var inputs []Input
	var outputs []Output

	for txid,outputIndexs := range container{
		for _,outputINdex := range outputIndexs{
			input := Input{[]byte(txid),outputINdex,from}
			inputs = append(inputs,input)
		}
	}

	output := Output{amount,to}
	outputs = append(outputs,output)
	if countTotal > amount{
		outputs = append(outputs,Output{countTotal-amount,from})
	}

	tx := transaction{nil,inputs,outputs}
	tx.SetTXID()
	return &tx
}


func (tx *transaction)IsCoinbase()bool{
	if len(tx.TXInputs) == 1{  //使用元素下标时一定要先判断是满足元素个数的！！
		if tx.TXInputs[0].Txid == nil && tx.TXInputs[0].ReferOutputIndex == -1{
			return true
		}
	}
	return false
}







