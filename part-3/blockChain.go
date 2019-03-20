package main

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {

	prevBlock := bc.Blocks[len(bc.Blocks)-1].Hash
	block := NewBlock(data, prevBlock)
	block.SetHash()
	bc.Blocks = append(bc.Blocks, block)

	return
}

func NewBlockChain() *BlockChain {

	genesisBlock := NewGenesisBlock()
	blcokChain := &BlockChain{
		[]*Block{
			genesisBlock,
		},
	}

	return blcokChain
}
