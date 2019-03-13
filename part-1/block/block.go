package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {

	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		nil,
	}
	block.SetHash()

	return block

}

func (b *Block) SetHash() {

	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}

func NewGenesisiBlock() *Block {

	return NewBlock("Genesis Block", []byte{})

}
