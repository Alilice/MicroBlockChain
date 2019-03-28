package main

import (
	"time"
	"crypto/sha256"
	"bytes"
)

type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Transactions   []*Transaction
	Nonce         int64
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
func NewBlock(prevBlockHash []byte, transactions []*Transaction) (block *Block) {

	block = &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Transactions:          transactions,
	}
	//block.setHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash[:]

	return
}

//HashTransactions 计算区块中所以交易的Hash
func (b *Block)HashTransactions() (txsHash []byte) {

	txHashes:=make([][]byte,0)

	for _,tx:=range b.Transactions{
		txHashes = append(txHashes,tx.TXHash)
	}
	txsHash=sha256.Sum256(bytes.Join(txHashes,[]byte{}))[:]

	return
}