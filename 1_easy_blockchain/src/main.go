package main

import (
	bc "github.com/DiDiDaDiDiDa/MicroBlockChain/1_easy_blockchain/src/blockchian"
	"fmt"
)

func main() {
	blockchain:=bc.NewGenesisBlockChain()

	data1:="第一个区块"
	data2:="第二个区块"
	data3:="第三个区块"

	blockchain.AddBlock(data1)
	blockchain.AddBlock(data2)
	blockchain.AddBlock(data3)
	for _,block:=range blockchain.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
