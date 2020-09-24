package pkg

type BlockChain struct {
	blks                []*Block
	difficulty          int
	pendingTransactions []*Transaction
	mineReward          int
}

func NewBlockChain(difficulty int) *BlockChain {
	chain := &BlockChain{
		difficulty:          difficulty,
		blks:                make([]*Block, 0),
		pendingTransactions: make([]*Transaction, 0),
		mineReward:          100,
	}
	chain.initChain()
	return chain
}

func (chain *BlockChain) initChain() {
	blk := NewBlock(make([]*Transaction, 0), "0")
	chain.blks = append(chain.blks, blk)
}

func (chain *BlockChain) AddTransaction(transaction *Transaction) {
	chain.pendingTransactions = append(chain.pendingTransactions, transaction)
}

func (chain *BlockChain) MineBlock(mineAddress string) {
	blk := NewBlock(chain.pendingTransactions, chain.LastBlock().hash)
	blk.MineBlock(chain.difficulty)

	chain.blks = append(chain.blks, blk)
	chain.pendingTransactions = []*Transaction{NewTransaction("", mineAddress, chain.mineReward)}
}

func (chain *BlockChain) GetBalance(address string) int {
	balance := 0
	for _, blk := range chain.blks {
		for _, transaction := range blk.transactions {
			if transaction.fromAddress == address {
				balance -= transaction.amount
			}
			if transaction.toAddress == address {
				balance += transaction.amount
			}
		}
	}
	return balance
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
