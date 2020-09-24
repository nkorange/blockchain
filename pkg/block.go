package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	transactions []*Transaction
	timestamp    int64
	hash         string
	preHash      string
	nonce        int
}

func NewBlock(transactions []*Transaction, preHash string) *Block {
	blk := &Block{
		transactions: transactions,
		timestamp:    time.Now().UnixNano(),
		preHash:      preHash,
	}
	blk.calculateHash()
	return blk
}

func (blk *Block) MineBlock(difficulty int) {
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix = prefix + "0"
	}
	for !strings.HasPrefix(blk.hash, prefix) {
		blk.nonce++
		blk.calculateHash()
	}
	log.Println("block mined", blk.hash)
}

func (blk *Block) calculateHash() {
	h := sha256.New()
	data, _ := json.Marshal(blk.transactions)
	h.Write(data)
	h.Write([]byte(blk.preHash))
	h.Write([]byte(strconv.Itoa(blk.nonce)))
	sum := h.Sum(nil)
	blk.hash = hex.EncodeToString(sum)
}

func (blk *Block) Hash() string {
	return blk.hash
}
