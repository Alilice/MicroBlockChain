package main

import (
	"time"
)

type Block struct {
	Timestamp int64
	PrevBlockHash []byte
	Hash []byte
	Data []byte
	Nonce int64
}

////setHash 计算区块链的Hash，这里只计算区块链的Timestamp+prevBlockHash+Data的hash
//func (b *Block) setHash() {
//
//	timestamp:=[]byte(strconv.FormatInt(b.Timestamp,10))
//	headers:=bytes.Join([][]byte{timestamp,b.PrevBlockHash,b.Data},[]byte{})
//	hash:=sha256.Sum256(headers)
//
//	b.Hash=hash[:]
//}

//NewBlock 创建新的Block区块
func NewBlock(prevBlockHash []byte ,data string ) (block *Block) {

	timestamp:=time.Now().Unix()
	block=&Block{
		Timestamp:timestamp,
		PrevBlockHash:prevBlockHash,
		Data:[]byte(data),
	}
	//block.setHash()
	pow:=NewProofOfWork(block)
	nonce,hash:=pow.Run()
	block.Nonce=nonce
	block.Hash=hash[:]

	return
}