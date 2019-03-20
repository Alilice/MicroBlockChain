package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
	"encoding/gob"
	"github.com/CyberMiles/go-ethereum/log"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {

	block := &Block{
		Timestamp:time.Now().Unix(),
		Data:[]byte(data),
		PrevBlockHash:prevBlockHash,
	}
	pow:=NewProofOfWork(block)
	nonce,hash:=pow.Run()
	block.Nonce=nonce
	block.Hash=hash[:]

	return block

}

func (b *Block) SetHash() {

	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}

func NewGenesisBlock() *Block {

	return NewBlock("Genesis Block", []byte{})

}

func (b *Block) Serialize() []byte {

	var result bytes.Buffer
	encoder:=gob.NewEncoder(&result)
	err:=encoder.Encode(b)
	if err!=nil {
		log.Error("序列化出现错误")
		return nil
	}
	return result.Bytes()
}