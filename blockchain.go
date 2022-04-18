package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// func NewBlock() *Block {
// 	return &Block{
// 		timestamp: time.Now().UnixNano(),
// 	}
// }

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block) //new performs &Block -> pointer to returned (*Block) above
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp            %d\n", b.timestamp)
	fmt.Printf("nonce	             %d\n", b.nonce)
	fmt.Printf("previous_hash        %s\n", b.previousHash)
	fmt.Printf("transactions         %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	b := NewBlock(0, "init hash")	//0 - nonce, init hash - for the first hash initialization
	b.Print()
}



//Running 1st time:
// ~/go/src/Blockchain$ go run blockchain.go 
// &{0 init hash 1650308612892056545 []} //log code defines the "timestamp", [] - defines transaction 