package wallet

type Wallet struct {
	balance int

	transactions []Transaction
}

type transType string

type Transaction struct {
	amount int
	transType
}

// List of Transaction of types.
const (
	Deposit  transType = "deposit"
	Withdraw transType = "withdraw"
)

// New initializes a new wallet with an initial balance.
func New(initialBal int) *Wallet {
	return &Wallet{balance: initialBal}
}

// Deposit adds the supplied amount to the Wallet's current balance.
func (w *Wallet) Deposit(amount int) {
	w.balance += amount

	w.transactions = append(w.transactions, Transaction{amount: amount, transType: Deposit})
}

// Withdraw substracts the supplied amount to the Wallet's current balance.
func (w *Wallet) Withdraw(amount int) {
	w.balance -= amount

	w.transactions = append(w.transactions, Transaction{amount: amount, transType: Withdraw})
}

// Balance returns the Wallet's current balance.
func (w *Wallet) Balance() int {
	return w.balance
}

// Transactions returns the Wallet's Transactions.
func (w *Wallet) Transactions() []Transaction {
	return w.transactions
}
