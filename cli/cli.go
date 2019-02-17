package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/yfsuperman/blockchain_learning_go/core"
)

// CLI is the general command-line operation type
type CLI struct {
	BC *core.Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage: blockchain_learning_go [OPTIONS] ARGS...")
	fmt.Println("")

	fmt.Println("options:")
	fmt.Println("  add-block    String    Input data for the new block")
	fmt.Println("  print-chain            Print the existing block chain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock(data)
	fmt.Println("Successfully added new block into the chain")
}

func (cli *CLI) printChain() {
	bci := cli.BC.Iterator()
	block := bci.Next()

	for len(block.PrevBlockHash) != 0 {
		fmt.Printf("Previous block hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Current block data: %s\n", block.Data)
		fmt.Printf("Current block hash: %x\n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		block = bci.Next()
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Construct block using this block data")
	addBlockDataShort := addBlockCmd.String("d", "", "Construct block using this block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		data := *addBlockData
		if *addBlockDataShort != "" {
			data = *addBlockDataShort
		}
		cli.addBlock(data)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}