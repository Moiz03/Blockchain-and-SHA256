package main

import (
	"crypto/sha256"
	"fmt"
)

// Block
type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {

	h := sha256.New() //calculating hash of strings
	h.Write([]byte(fmt.Sprintf("%v", inputBlock.transactions)))
	transaction := fmt.Sprintf("%x", h.Sum(nil))

	h = sha256.New() //calculating hash of previous pointer
	h.Write([]byte(fmt.Sprintf("%v", inputBlock.prevPointer)))
	prevPointer := fmt.Sprintf("%x", h.Sum(nil))

	h = sha256.New() //calculating hash of previous hash
	h.Write([]byte(fmt.Sprintf("%v", inputBlock.prevHash)))
	prevHash := fmt.Sprintf("%x", h.Sum(nil))

	//current hash is made by adding hashes of all the element of block 
	var hash string
	hash = transaction + prevPointer + prevHash
	return hash //returning hash
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {

	//insert new block and return head pointer
	var newChainHead Block

	if chainHead == nil { //checking if gensis block
		//if yes then make previous pointer null and previous hash empty
		newChainHead.prevPointer = nil
		newChainHead.prevHash = ""
	} else { 
		//if no
		newChainHead.prevPointer = chainHead
		newChainHead.prevHash = chainHead.currentHash
	}

	newChainHead.transactions = transactionsToInsert
	newChainHead.currentHash = CalculateHash(&newChainHead) //calculating hash

	return &newChainHead //returning new chainHead
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	//change transaction data inside block
	for index := range chainHead.transactions {
		if chainHead.transactions[index] == oldTrans { //checking if current index string is equal to oldTrans
			chainHead.transactions[index] = newTrans //if yes then replace that string with newTrans
		}
	}

	chainHead.currentHash = CalculateHash(chainHead) //recalculate current hash and replace it with old current hash
}

func ListBlocks(chainHead *Block) {

	//dispaly the data(transaction) inside all blocks
	var count int
	count = 0
	for index := chainHead; index != nil; index = index.prevPointer {
		count++ 
		fmt.Printf("Transactions in Block # %d from last\n", count)
		for str := range index.transactions { //printing all the strings in block
			println(index.transactions[str])
		}
		print("\n")
	}
}

func VerifyChain(chainHead *Block) {

	//check whether "Block chain is compromised" or "Block chain is unchanged"

	flag := true

	for index := chainHead; index.prevPointer != nil && flag == true; index = index.prevPointer {
		if index.prevHash != index.prevPointer.currentHash { 
			//if previous hash in current block is not equal to current hash of previous block
			//then blockchain is compromised
			flag = false
		}
	}

	if flag == false {
		println("Block chain is compromised\n")
	} else {
		println("Block chain is unchanged\n")
	}

}

func main() {

	//inserting values
	gensisBlock := InsertBlock([]string{"Transaction 1", "Transaction 2", "Transaction 3"}, nil)
	firstBlock := InsertBlock([]string{"Transaction 4", "Transaction 5", "Transaction 6"}, gensisBlock)
	secondBlock := InsertBlock([]string{"Transaction 7", "Transaction 8", "Transaction 9"}, firstBlock)
	thirdBlock := InsertBlock([]string{"Transaction 10", "Transaction 11", "Transaction 12"}, secondBlock)

	//displaying all blocks transactions
	ListBlocks(thirdBlock)

	//verifying the blockchain
	VerifyChain(thirdBlock)

	//changing transaction in block of blockchain
	ChangeBlock("Transaction 8", "Transaction 99", secondBlock)

	//displaying all blocks transactions
	ListBlocks(thirdBlock)
	
	//verifying the blockchain
	VerifyChain(thirdBlock)
}
