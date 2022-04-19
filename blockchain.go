package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

//Block section
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
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
	fmt.Printf("previous_hash        %x\n", b.previousHash)
	fmt.Printf("transactions         %s\n", b.transactions)
}

//End Block section

//Hash section
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256(([]byte(m)))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp int64 `json:"timestamp"`			//in `` we define desired jason output names
		Nonce int 		`json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp: b.timestamp,
		Nonce: b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

//End Hash section

//Blockchain section
//Blockchain struct created
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

//CreateBlock method created
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain) - 1]
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

	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
}



