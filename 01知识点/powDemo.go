package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	src := "hello blockchain"

	for i := 0; i < 10; i++ {
		hash := sha256.Sum256([]byte(src + string(i)))
		s := hex.EncodeToString(hash[:])
		fmt.Println(s)
	}
}
