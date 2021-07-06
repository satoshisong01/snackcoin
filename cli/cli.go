package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/sks8982/snackcoin/explorer"
	"github.com/sks8982/snackcoin/rest"
)

//커맨드라인 플래그 선언
func usage() {
	fmt.Printf("Welcome to 윤원 코인\n\n")
	fmt.Printf("커맨드를 입력해 주세요\n\n")
	fmt.Printf("-port: Set PORT of the Server\n")
	fmt.Printf("-mode: html 과 rest api중 골라주세요\n\n")
	os.Exit(0) //종료
}

func Start() {
	//os.Args는 커맨드라인 인자를 그대로 접근하는 방법을 제공 (1은 프로그램의 인자를 가짐)
	if len(os.Args) == 1 {
		usage()
	}

	//포트int
	port := flag.Int("port", 5000, "Set port of the Server")

	mode := flag.String("mode", "rest", "Choose 'html' or 'rest api'")
	//커맨드라인 파싱
	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	//플래그 선언 함수에 변수의 포인터를 전달
	//실제 옵션값을 얻기 위해서 포인터 역참조
	fmt.Println(*port, *mode)
}
