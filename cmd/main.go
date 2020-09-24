package main

import "github.com/nkorange/blockchain/pkg"

func main() {
	chain := pkg.NewBlockChain(5)
	blk := pkg.NewBlock([]byte("data1"), chain.LastBlock().Hash())
	blk.MineBlock(chain.Difficulty())
}
