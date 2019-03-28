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

	defer blockchain.db.Close()

	cli := &CLI{blockchain}
	cli.Run()

}
