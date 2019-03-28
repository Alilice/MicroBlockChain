package main

import (
	"fmt"
)

func main() {
	blockchain:=NewGenesisBlockChain()

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
		pow:=NewProofOfWork(block)
		fmt.Printf("POW:%t\n",pow.Validate())
		fmt.Println()
	}

}
