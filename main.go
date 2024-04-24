package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
)

// BlockChain represents a blockchain containing multiple blocks.
type BlockChain struct {
	blocks []*Block // Array of blocks
}

// Block represents a block in the blockchain.
type Block struct {
	Hash     []byte // Hash of the current block
	Data     []byte // Data stored in the block
	PrevHash []byte // Hash of the previous block
}

// DeriveHash calculates the hash of a block based on its data and the hash of the previous block.
func (b *Block) DeriveHash() {
	// Concatenate the data and previous hash to create the input for hash calculation
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Calculate the hash using SHA-256 algorithm
	hash := sha256.Sum256(info)
	// Assign the calculated hash to the block's Hash field
	b.Hash = hash[:]
}

// CreateBlock creates a new block with the given data and previous hash, and derives its hash.
func CreateBlock(data string, prevHash []byte) *Block {
	// Create a new block instance with empty hash, provided data, and previous hash
	block := &Block{[]byte{}, []byte(data), prevHash}
	// Derive and set the hash for the block
	block.DeriveHash()
	// Return the created block
	return block
}

// AddBlock adds a new block to the blockchain with the given data.
func (chain *BlockChain) AddBlock(data string) {
	// Get the previous block from the blockchain
	prevBlock := chain.blocks[len(chain.blocks)-1]
	// Create a new block with the provided data and previous block's hash
	new := CreateBlock(data, prevBlock.Hash)
	// Append the new block to the blockchain
	chain.blocks = append(chain.blocks, new)
}

// Genesis creates the genesis block of the blockchain.
func Genesis() *Block {
	// Create the genesis block with default data "Genesis" and empty previous hash
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initializes a new blockchain with the genesis block.
func InitBlockChain() *BlockChain {
	// Create a new blockchain instance with the genesis block as its only block
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	// Initialize the blockchain
	chain := InitBlockChain()

	// Add blocks to the blockchain
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// Print information about each block in the blockchain
	for _, block := range chain.blocks {
		// Print the previous block's hash
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		// Print the data stored in the block
		fmt.Printf("Data in Block: %s\n", block.Data)
		// Print the hash of the current block
		fmt.Printf("Hash: %x\n", block.Hash)
		// Print a separator for readability
		fmt.Println(strings.Repeat("-", 25), "chain", strings.Repeat("-", 25))
	}
}
