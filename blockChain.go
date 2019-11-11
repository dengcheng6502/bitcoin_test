package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

const dbFile = "blockChain.db"
const blockBucket = "bucket"
const lastHashkey = "lastkey"
const genesisInfo = "genesis"

type BlockChain struct {
	//数据库操作句柄
	db   *bolt.DB
	tail []byte
}

func IsDBExist() bool {
	_, err := os.Stat(dbFile)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func InitBlockChain(address string) *BlockChain {
	if IsDBExist() {
		fmt.Println("blockchain exist already,no need to create!")
		os.Exit(1)

	}

	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("NewBlockChain1", err)

	var lastHash []byte

	db.Update(func(tx *bolt.Tx) error {

		coinbase := NewCoinbaseTx(address, genesisInfo)
		genesis := NewGenesisBlock(coinbase)
		bucket, err := tx.CreateBucket([]byte(blockBucket))
		CheckErr("NewBlockChain2", err)
		err = bucket.Put(genesis.Hash, genesis.Serialize())
		CheckErr("NewBlockChain3", err)
		err = bucket.Put([]byte(lastHashkey), genesis.Hash)
		CheckErr("NewBlockChain4", err)
		lastHash = genesis.Hash

		return nil
	})
	return &BlockChain{db, lastHash}

}

func GetBlockChainHandler() *BlockChain {

	if !IsDBExist() {
		fmt.Println("Please create blockchain first!")
		os.Exit(1)

	}

	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("GetBlockChainHandler1", err)

	var lastHash []byte

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			lastHash = bucket.Get([]byte(lastHashkey))
		}

		return nil
	})
	return &BlockChain{db, lastHash}
}
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	var prevBlockHash []byte
	bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		prevBlockHash = bucket.Get([]byte(lastHashkey))
		return nil
	})
	block := NewBlock(txs, prevBlockHash)
	err := bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		err := bucket.Put(block.Hash, block.Serialize())

		CheckErr("AddBlock1", err)
		err = bucket.Put([]byte(lastHashkey), block.Hash)
		CheckErr("AddBlock2", err)
		bc.tail = block.Hash
		return nil
	})
	CheckErr("AddBlock3", err)

}

type BlockChainIterator struct {
	currHash []byte
	db       *bolt.DB
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{currHash: bc.tail, db: bc.db}
}

func (it *BlockChainIterator) Next() (block *Block) {
	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			return nil
		}

		data := bucket.Get(it.currHash)
		block = Deserialize(data)
		it.currHash = block.PrevBlockHash

		return nil
	})
	CheckErr("Next()", err)
	return
}
