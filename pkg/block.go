package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	data      []byte
	timestamp int64
	hash      string
	preHash   string
}

func NewBlock(data []byte, preHash string) *Block {
	blk := &Block{
		data:      data,
		timestamp: time.Now().UnixNano(),
		preHash:   preHash,
	}
	blk.calculateHash()
	return blk
}

func (blk *Block) calculateHash() {
	h := sha256.New()
	h.Write(append(blk.data, []byte(blk.preHash)...))
	sum := h.Sum(nil)
	blk.hash = hex.EncodeToString(sum)
}
