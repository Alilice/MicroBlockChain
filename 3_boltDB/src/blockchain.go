package main

import (
	"github.com/boltdb/bolt"
	"log"
)

var (
	dbfile       = "./3_boltDB/blockchain.db"
	blocksBucket = []byte("BlocksBucket")
	lastBlockKey = []byte("l")
)

type BlockChain struct {
	//Blocks []*Block
	tip []byte
	db  *bolt.DB
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

//AddBlock 向区块链上增加区块
func (bc *BlockChain) AddBlock(data string) (err error) {

	//拿到最新的区块链的hash
	var lastBlockHash []byte
	err = bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(blocksBucket)
		lastBlockHash = b.Get(lastBlockKey)
		return nil
	})
	if err != nil {
		log.Println("查询数据库出错！")
		return
	}

	//组装区块
	newBlock := NewBlock(lastBlockHash, data)

	//更新数据库
	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(blocksBucket)
		errs := make([]error, 2)
		errs[0] = bucket.Put(newBlock.Hash, Serialize(newBlock))
		errs[1] = bucket.Put(lastBlockKey, newBlock.Hash)
		for _, err := range errs {
			if err != nil {
				log.Println("插入数据库出错！")
				return err
			}
		}
		bc.tip = newBlock.Hash
		return nil
	})

	if err != nil {
		log.Println("更新数据库出错！")
	}

	return
}

//NewGenesisBlockChain 创建创世链
func NewGenesisBlock() (b *Block) {

	b = NewBlock(nil, "Genesis Block!")

	return
}

//NewBlockChain 初始化BlockChain
func NewBlockChain() (bc *BlockChain, err error) {

	////检查文件是否存在，不存在创建一个
	//err=CreateFileIfNotExist(dbfile)
	//if err != nil {
	//	return
	//}

	var tip []byte
	db, err := bolt.Open(dbfile, 0600, nil)
	if err != nil {
		log.Println("打开数据库文件失败！")
		return
	}

	err = db.Update(func(tx *bolt.Tx) error {
		//检查数据库文件是否存在，不存在就创建一个
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			//创建数据库
			bucket, err := tx.CreateBucket(blocksBucket)
			if err != nil {
				log.Println("创建bucket失败！")
				return err
			}

			//创建创世块并插入数据库
			genesis := NewGenesisBlock()
			errs := make([]error, 2)
			errs[0] = bucket.Put(genesis.Hash, Serialize(genesis))
			errs[1] = bucket.Put(lastBlockKey, genesis.Hash)
			for _, err := range errs {
				if err != nil {
					log.Println("插入数据库出错！")
					return err
				}
			}
			tip = genesis.Hash
		} else {
			tip = b.Get(lastBlockKey)
		}
		return nil
	})
	if err != nil {
		log.Println("数据库操作出错！")
		return
	}

	//组装BlockChain
	bc = &BlockChain{
		tip: tip,
		db:  db,
	}

	return
}

func (bc *BlockChain) Iterator() (bci *BlockChainIterator) {
	bci = &BlockChainIterator{
		bc.tip,
		bc.db,
	}
	return
}

func (bci *BlockChainIterator) Next() (b *Block) {
	b = new(Block)
	bci.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(blocksBucket)
		blockBytes := bucket.Get(bci.currentHash)
		Deserialize(blockBytes, b)
		return nil
	})
	bci.currentHash = b.PrevBlockHash
	return
}
