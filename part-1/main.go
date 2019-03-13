package main

import (
	"fmt"
	"github.com/DiDiDaDiDiDa/MicroBlockChain/part-1/block"
)

func main() {

	bc := block.NewBlockChain()
	bc.AddBlock("第二个区块交易")
	bc.AddBlock("第三个区块交易")

	fmt.Printf("%v\n", bc)

	for k, v := range bc.Blocks {
		fmt.Printf("第%d个区块：\n", k)
		fmt.Println("----------------")
		fmt.Printf("前一个区块的Hash：%x\n", v.PrevBlockHash)
		fmt.Printf("区块的数据：%s\n", v.Data)
		fmt.Printf("区块的Hash：%x\n", v.Hash)
		fmt.Println("")
	}

	return
}
