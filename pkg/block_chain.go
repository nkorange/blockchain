package pkg

type BlockChain struct {
	blks []*Block
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
