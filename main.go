package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash)) //byte값 가져오기
	hexHash := fmt.Sprintf("%x", hash)                                       //string화
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}


func sumAndMult(a, b int) (int, int) {
	return (a + b), (a * b)
}