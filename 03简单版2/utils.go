package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func IntToByte(num uint64)[]byte{
	var buf bytes.Buffer

	err := binary.Write(&buf, binary.BigEndian, num)
	checkErr(err)

	return buf.Bytes()
}


func checkErr(err error){
	if err != nil{
		fmt.Println(" err occur:",err)
		os.Exit(1)
	}
}