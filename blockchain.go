package main

import (
	"fmt"
	"log"
	"strings"
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

//Print method added to Block
func (b *Block) Print() {
	fmt.Printf("timestamp            %d\n", b.timestamp)
	fmt.Printf("nonce	             %d\n", b.nonce)
	fmt.Printf("previous_hash        %s\n", b.previousHash)
	fmt.Printf("transactions         %s\n", b.transactions)
}

//Blockchain struct created
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

//CreateBlock method created
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

//Print method added to Blockchain
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}


//Running before Print method added
// ~/go/src/Blockchain(main)$ go run blockchain.go 
// &{[] [0xc0000b8040]}						     //[] - empty transactionPool, [0xc0000b8040] - chain/new Blockchain


// Running after Print method added
// ~/go/src/Blockchain(main)$ go run blockchain.go 
// Chain 0 
// timestamp            1650312388953139233
// nonce                0
// previous_hash        init hash
// transactions         []


//Running after refactoring Print method
// ~/go/src/Blockchain(main)$ go run blockchain.go 
// ========================= Chain 0 =========================
// timestamp            1650312974920236803
// nonce                0
// previous_hash        init hash
// transactions         []
// *************************
// ========================= Chain 0 =========================
// timestamp            1650312974920236803
// nonce                0
// previous_hash        init hash
// transactions         []
// ========================= Chain 1 =========================
// timestamp            1650312974920319322
// nonce                5
// previous_hash        hash 1
// transactions         []
// *************************
// ========================= Chain 0 =========================
// timestamp            1650312974920236803
// nonce                0
// previous_hash        init hash
// transactions         []
// ========================= Chain 1 =========================
// timestamp            1650312974920319322
// nonce                5
// previous_hash        hash 1
// transactions         []
// ========================= Chain 2 =========================
// timestamp            1650312974920362485
// nonce                2
// previous_hash        hash 2
// transactions         []
// *************************