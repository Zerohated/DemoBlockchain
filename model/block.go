package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Block struct {
	Index      int
	Timestamp  string
	Info       string
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}
type Message struct {
	Info string
}

const Difficulty = 1

var (
	Blockchain []Block
	Mutex      = &sync.Mutex{}
)

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Info + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func IsHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
func GenerateBlock(oldBlock Block, info string) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Info = info
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = Difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !IsHashValid(CalculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(CalculateHash(newBlock), " do more work!")
			time.Sleep(time.Second / 500)
			continue
		} else {
			fmt.Println(CalculateHash(newBlock), " work done!")
			newBlock.Hash = CalculateHash(newBlock)
			break
		}

	}
	return newBlock
}

func IsBlockValid(newBlock Block, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
