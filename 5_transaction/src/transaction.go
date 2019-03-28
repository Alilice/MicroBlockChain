package main

import (
	"fmt"
	"log"
	"encoding/hex"
)

var subsidy int64 = 50

type TXInput struct {
	TXHash      []byte
	OutputID  int64
	ScriptSig string
}

type TXOutput struct {
	Value        int64
	ScriptPubKey string
}

type Transaction struct {
	TXHash  []byte
	Inputs  []*TXInput
	Outputs []*TXOutput
}

//NewUTXOTransaction 创建新的交易
func NewUTXOTransaction(from, to string, amount int64, bc *BlockChain) (tx *Transaction) {

	inputs :=make([]*TXInput,0)
	outputs :=make([]*TXOutput,0)

	//首先找到自己可用于交易的coin
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	//检查上面所得的是否足够用于交易
	if acc < amount {
		log.Panic("ERROR: Not enough funds")
		return
	}
	//如果足够，组装交易
	//组装input
	for tx,outs:=range validOutputs{
		txHash,err:=hex.DecodeString(tx)
		if err != nil {
			log.Printf("交易哈希解析成[]byte出错！")
		}
		for _,outID:=range outs{
			input:=&TXInput{
				TXHash:txHash,
				OutputID:outID,
				ScriptSig:from,
			}

			inputs=append(inputs,input)
		}
	}

	//组装output
	output:=&TXOutput{
		Value:amount,
		ScriptPubKey:to,
	}
	outputs = append(outputs,output)
	//需要找零的情况
	if acc > amount{
		changeOutput:=&TXOutput{
			Value:acc-amount,
			ScriptPubKey:from,
		}
		outputs = append(outputs,changeOutput)
	}

	tx=&Transaction{
		Inputs:inputs,
		Outputs:outputs,
	}
	tx.SetHash()

	return
}

//NewCoinBaseTX 创建一个 coinbase 交易
func NewCoinBaseTX(to, data string) (tx *Transaction) {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txin := &TXInput{[]byte{}, -1, data}
	txout := &TXOutput{subsidy, to}
	tx = &Transaction{nil, []*TXInput{txin}, []*TXOutput{txout}}
	tx.SetHash()

	return
}

//SetID 计算交易的Hash
func (tx *Transaction) SetHash() {

}