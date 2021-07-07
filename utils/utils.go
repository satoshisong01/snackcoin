package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//byte 로바꿔주는 인터페이스 (일반 문자열 -> 인코딩 문자열)
func Tobytes(i interface{}) []byte { //byte로 encoding한뒤 그 bytes를 읽음
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

//interface를 받아 data를 interface안에 decode함 (인코딩 문자열 -> 일반 문자열)
func FromBytes(i interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}
