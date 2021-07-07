package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/sks8982/snackcoin/db"
	"github.com/sks8982/snackcoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHas,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() { // block을 저장하기 위해만든 SaveBlock함수 호출
	db.SaveBlock(b.Hash, utils.Tobytes(b)) //블록의 hash key로 사용해서 찾음(bytes 만 저장하고 받음)
}

var ErrNotFound = errors.New("블록을 찾을 수 없습니다.")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) { //hash를 받고  Block포인터 혹은 error를 리턴
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}         //빈블록 생성
	block.restore(blockBytes) //빈블록에 이전 블록 복원
	return block, nil
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return &block
}
