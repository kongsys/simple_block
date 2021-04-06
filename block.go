package main

import (
	"time"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	TimeStamp	time.Time
	TransActions	[]string
	PrevHash	[]byte
	Hash		[]byte
}

func NewBlock(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()
	return &Block {
		TimeStamp: currentTime,
		TransActions: transactions,
		PrevHash: prevHash,
		Hash: NewHash(currentTime, transactions, prevHash),
	}
}

func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func printBlockInfo(block *Block) {
	fmt.Printf("\ttime: %s\n", block.TimeStamp.String())
	fmt.Printf("\tprevHash: %x\n", block.PrevHash)
	fmt.Printf("\tHash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.TransActions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}

func main() {
	genesisTransactions := []string{"Bob sent will 50 bitcoin", "Will sent Bob 30 bitcoin"}
	genesisBlock := NewBlock(genesisTransactions, []byte{})
	fmt.Println("--- First Block ---")
	printBlockInfo(genesisBlock)

	block2Transactions := []string{"Alice sent Bob 30 bitcoin"}
	block2 := NewBlock(block2Transactions, genesisBlock.Hash)
	fmt.Println("--- Second Block ---")
	printBlockInfo(block2)

	block3Transactions := []string{"Will sent Bob 45 bitcoin", "Bob sent 10 bitcoin"}
	block3 := NewBlock(block3Transactions, block2.Hash)
	fmt.Println("--- Second Block ---")
	printBlockInfo(block3)

}
