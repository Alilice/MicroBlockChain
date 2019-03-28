package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	bc *BlockChain
}

var (
	addblock   = "addblock"
	printchain = "printchain"
	usage      = `
		Usage:
  			addblock -data BLOCK_DATA    add a block to the blockchain
  			printchain                   print all the blocks of the blockchain
		`
)

//Run 运行cli
func (cli *CLI) Run() {

	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet(addblock, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(printchain, flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case addblock:
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Print("参数解析出错！")
			return
		}
	case printchain:
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Println("参数解析出错！")
			return
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		err := cli.addBlock(*addBlockData)
		if err != nil {
			log.Println("加入区块链出错！")
			os.Exit(1)
		}
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
	return
}

//printUsage 对不存在的命令进行处理
func (cli *CLI) printUsage() {
	log.Println(usage)
	return
}

//validateArgs 验证参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
	return
}

//printChain 打印区块链的区块信息
func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		b := bci.Next()
		fmt.Printf("BlockInfo:\n Data:%s,\n PrevBlockHash:%x,\n Hash:%x,\n Timestamp:%d,\n Nonce:%d", b.Data, b.PrevBlockHash, b.Hash, b.Timestamp, b.Nonce)

		if len(b.PrevBlockHash) == 0 {
			break
		}
	}
	return
}

//validateArgs 验证参数
func (cli *CLI) addBlock(data string) (err error) {
	err = cli.bc.AddBlock(data)
	return
}
