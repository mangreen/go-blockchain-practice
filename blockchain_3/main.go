package main

import (
	"blockchain_3/src"
)

func main() {
	bc := src.NewBlockchain()
	defer bc.DB.Close()

	cli := src.CLI{BC: bc}
	cli.Run()
}
