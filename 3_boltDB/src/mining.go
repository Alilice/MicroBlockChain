package main

import (
	"math/big"

	"bytes"
	"crypto/sha256"
	"math"
)

const targetBits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//NewProofOfWork 初始化ProofOfWork
func NewProofOfWork(b *Block) (pow *ProofOfWork) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow = &ProofOfWork{
		block:  b,
		target: target,
	}
	return
}

//prepareData 准备要计算的数据
func (pow *ProofOfWork) prepareData(nonce int64) (data []byte) {

	data = bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			Int2Byte(pow.block.Timestamp),
			Int2Byte(int64(targetBits)),
			Int2Byte(nonce),
		}, []byte{})

	return
}

//Run 找到合适的nonce
func (pow *ProofOfWork) Run() (nonce int64, hash [32]byte) {
	hashInt := new(big.Int)

	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return
}

func (pow *ProofOfWork) Validate() (isValid bool) {

	hashInt := new(big.Int)
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid = hashInt.Cmp(pow.target) == -1

	return
}
