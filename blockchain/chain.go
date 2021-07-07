package blockchain

import (
	"sync"

	"github.com/sks8982/snackcoin/db"
	"github.com/sks8982/snackcoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.Tobytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash //업데이트
	b.Height = block.Height   //업데이트
}

//전체블록 찾기
func (b *blockchain) Blocks() []*Block {
	var blocks []*Block        //블록포인터를 만들고
	hashCursor := b.NewestHash //찾을 hash인 hashCursor를 만들고 (처음에는 NewestHash를찾음)
	for {
		block, _ := FindBlock(hashCursor) //블록체인의 NewestHash로 블록을 찾음
		blocks = append(blocks, block)    //블록을 찾으면 slice에 넣어주고
		if block.PrevHash != "" {         //PrevHash가 빈 string이 아니면 hashCursor를 찾은 블록의 prevHash로 바꿈
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

//접근시 GetBlockchain 실행함 (처음 시작시)
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}        //시작시 빈값에 0번째
			checkpoint := db.Checkpoint() //checkpoint에 data가 있는지 확인
			if checkpoint == nil {
				b.AddBlock(("Genesis")) //add -> createblock persist() -> byte변환 -> 생성 -> blockchain 업데이트
			} else {
				b.restore(checkpoint)
			}
		})
	}
	return b
}
