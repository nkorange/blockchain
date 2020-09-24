package pkg

type Transaction struct {
	fromAddress string
	toAddress   string
	amount      int
}

func NewTransaction(from, to string, amount int) *Transaction {
	return &Transaction{
		fromAddress: from,
		toAddress:   to,
		amount:      amount,
	}
}
