package blockchian

type BlockChain struct {
	Blocks []*Block
}

//AddBlock 向区块链上增加区块
func (bc *BlockChain) AddBlock(data string) {

	prevBlockHash := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(prevBlockHash.Hash, data)
	bc.Blocks = append(bc.Blocks, block)

	return
}

//NewGenesisBlockChain 创建创世链
func NewGenesisBlockChain() (bc *BlockChain) {

	block := NewBlock(nil, "Genesis Block!")
	bc = new(BlockChain)
	bc.Blocks = append(bc.Blocks, block)

	return
}
