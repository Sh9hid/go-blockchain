package main


import (
	"fmt"
)
func main() {
	bc:= New Blockchain()

	bc.AddBlock("Send 1 BTC to Sh9")
	bc.AddBlock("Send 2 more BTC to Sh9")

	for _, block = range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}