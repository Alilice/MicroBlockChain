package main

import (
	"fmt"
)

func main() {
	blockchain, err := NewBlockChain()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	data1 := "第一个区块"

	blockchain.AddBlock(data1)

	bci := blockchain.Iterator()
	b := bci.Next()

	fmt.Printf("BlockInfo:\n PrevBlockHash:%x,\n Hash:%x,\n Timestamp:%d,\n Data:%s,\n Nonce:%d", b.PrevBlockHash, b.Hash, b.Timestamp, b.Data, b.Nonce)

}
