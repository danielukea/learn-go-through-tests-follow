package pointers

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// w is a copy of the struct that we are pulling from * points
func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}
