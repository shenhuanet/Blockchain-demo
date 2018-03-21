package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("A向B发出100比特币.")
	bc.AddBlock("A向B发出200比特币.")

	for _, block := range bc.blocks {
		fmt.Printf("上一区块哈希:%x\n", block.PrevBlockHash)
		fmt.Printf("当前区块数据:%s\n", block.Data)
		fmt.Printf("当前区块哈希:%x\n\n", block.Hash)
	}
}
