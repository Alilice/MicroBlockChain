package main

import (
	"fmt"
	"strconv"
)

func main() {

	bc := NewBlockChain()
	bc.AddBlock("第二个区块交易")
	bc.AddBlock("第三个区块交易")

	fmt.Printf("%v\n", bc)

	for k, v := range bc.Blocks {
		fmt.Printf("第%d个区块：\n", k)
		fmt.Println("----------------")
		fmt.Printf("前一个区块的Hash：%x\n", v.PrevBlockHash)
		fmt.Printf("区块的数据：%s\n", v.Data)
		fmt.Printf("区块的Hash：%x\n", v.Hash)
		pow := NewProofOfWork(v)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("")
	}

	return
}
