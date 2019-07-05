package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func IntToByte(num int64)[]byte{
	var buf bytes.Buffer

	err := binary.Write(&buf, binary.BigEndian, num)
	checkErr(err,"binary.Write")

	return buf.Bytes()
}


func checkErr(err error,info string){
	if err != nil{
		fmt.Println(info," err occur:",err)
		os.Exit(1)
	}
}

//big.Int里直接有一个函数 setBytes 可以把byte转化为int
//func ByteToInt(data []byte) int64{
//	var buf bytes.Buffer
//
//	err := binary.Write(&buf, binary.BigEndian, data)
//	checkErr(err)
//
//	return int64(buf.Cap())
//}