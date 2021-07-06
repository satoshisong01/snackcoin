package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sks8982/snackcoin/blockchain"
	"github.com/sks8982/snackcoin/utils"
)

var port string

//const port string = ":4000"

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"` //json:000 으로 대문자를 소문자로 바꿀수있음
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` //omitempty 로 비어있으면 출력하지않고 가려준다
}

type addBlockBody struct {
	Message string
}

type errorRespone struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{ //Marshal은 JSON으로 encoding한 인터페이스를 리턴 (메모리형식으로 저장된 객체를 저장/송신 할수있도록 변환)
		{ //유저에게 보여줄 내용
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A block",
			Payload:     "data:string",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody                                  //변수, 구조체 두개
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody)) //쓸때는 Encode 읽어올때는 Decode / utils.handleErr 에러처리
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)      //블록체인 추가
		rw.WriteHeader(http.StatusCreated)                             //status code보내기 201
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id) //go언어 형변환
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorRespone{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

//텍스트로 인식하는걸 json으로 변경 (go에서 Middleware는 http.Handler를 받고 return하는 함수)
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	router := mux.NewRouter() //gorilla/mux 사용
	port = fmt.Sprintf(":%d", aPort)
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET") //gorilla/mux를 사용하여 아이디에 조건을 넣을수 있음
	fmt.Printf("http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
