package pkg

type BlockChain struct {
	blks       []*Block
	difficulty int
}

func NewBlockChain(difficulty int) *BlockChain {
	chain := &BlockChain{
		difficulty: difficulty,
		blks:       make([]*Block, 1),
	}
	chain.initChain()
	return chain
}

func (chain *BlockChain) initChain() {
	blk := NewBlock([]byte("0"), "0")
	chain.blks = append(chain.blks, blk)
}

func (chain *BlockChain) AddBlock(block *Block) {
	if len(chain.blks) == 0 {
		chain.blks = append(chain.blks, block)
		return
	}
	lastBlk := chain.blks[len(chain.blks)-1]
	block.preHash = lastBlk.hash
	block.calculateHash()
	chain.blks = append(chain.blks, block)
}

func (chain *BlockChain) LastBlock() *Block {
	return chain.blks[len(chain.blks)-1]
}

func (chain *BlockChain) Difficulty() int {
	return chain.difficulty
}
