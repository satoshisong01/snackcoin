package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to 윤원 코인\n\n")
	fmt.Printf("커맨드를 입력해 주세요\n\n")
	fmt.Printf("explorer: Start HTML Explorer\n")
	fmt.Printf("rest: Start REST API (recommended)\n\n")
	os.Exit(0) //종료
}

func main() {

	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of the server(default 5000)")

	fmt.Println(os.Args[2:])

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start server")
	}
}
