package main

import (
	"github.com/sks8982/snackcoin/cli"
	"github.com/sks8982/snackcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
