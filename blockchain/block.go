package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/sks8982/snackcoin/db"
	"github.com/sks8982/snackcoin/utils"
)

const difficulty int = 2 //hash 앞에 0의 갯수

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHas,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
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

//채굴
func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		blockAsString := fmt.Sprint(b)                                  //블록 string생성
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString))) //string을 hash로바꾸고 hash를 다시 16진수로 바꿈
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsString, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) { //target이랑 동일한 형태로 시작하는지 확인
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}
	block.mine()
	block.persist()
	return &block
}
