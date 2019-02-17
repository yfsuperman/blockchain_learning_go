package main

import (
	"github.com/yfsuperman/blockchain_learning_go/core"
	"github.com/yfsuperman/blockchain_learning_go/cli"
)

func main() {
	bc := core.NewBlockchain()

	defer bc.Db.Close()

	cli := cli.CLI{bc}
	cli.Run()
}