package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	src := []byte("hello base64")
	s := base64.StdEncoding.EncodeToString(src)
	fmt.Println(s)

	bytes, e := base64.StdEncoding.DecodeString(s)
	if e != nil{
		fmt.Println(e)
		return
	}
	fmt.Println(string(bytes))
}
