package src

import "github.com/boltdb/bolt"

const dbFile = "blockChain.db"
const blockBucket = "bucket"
const lastHashkey = "lastkey"

type BlockChain struct {
	//数据库操作句柄
	db   *bolt.DB
	tail []byte
}

func NewBlockChain() *BlockChain {

	db, err := bolt.Open(dbFile, 0x0600, nil)
	CheckErr("NewBlockChain1", err)

	var lastHash []byte

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			lastHash = bucket.Get([]byte(lastHashkey))
		} else {
			genesis := NewGenesisBlock()
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			CheckErr("NewBlockChain2", err)
			err = bucket.Put(genesis.Hash, genesis.Serialize()) //TODO 写一种方法使得区块的数据序列化后返回一个[]byte
			CheckErr("NewBlockChain3", err)
			err = bucket.Put([]byte(lastHashkey), genesis.Serialize())
			CheckErr("NewBlockChain4", err)
			lastHash = genesis.Hash
		}

		return nil
	})
	return &BlockChain{db, lastHash}

}

func (bc *BlockChain) AddBlock(data string) {
	prevBlockHash := bc.blocks[len(bc.blocks)-1].Hash
	block := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks, block)
}
