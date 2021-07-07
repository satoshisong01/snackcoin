package db

import (
	"github.com/boltdb/bolt"
	"github.com/sks8982/snackcoin/utils"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
	checkpoint   = "checkpoint"
)

//bolt db프레임워크 키와 벨류로만 이루워져있음
var db *bolt.DB

//내용을 담을 bucket 선언
func DB() *bolt.DB {
	if db == nil {
		dbPoint, err := bolt.Open(dbName, 0600, nil) //블록체인db에 작성 0600(write)
		db = dbPoint
		utils.HandleErr(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket)) //버킷 선언
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket)) //없으면 에러
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func Close() {
	DB().Close()
}

func SaveBlock(hash string, data []byte) { //hash Key와 데이터
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveCheckpoint(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func Block(hash string) []byte { //hash를받아 blockBucket에서 key를 찾음
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
