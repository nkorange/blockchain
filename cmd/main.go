package main

import (
	"github.com/nkorange/blockchain/pkg"
	"log"
)

func main() {
	chain := pkg.NewBlockChain(5)
	chain.AddTransaction(pkg.NewTransaction("addr1", "addr2", 100))
	chain.AddTransaction(pkg.NewTransaction("addr2", "addr1", 50))

	chain.MineBlock("addr3")
	log.Println(chain.GetBalance("addr3"))

	chain.MineBlock("addr3")
	log.Println(chain.GetBalance("addr3"))
}
