package wallet

import (
	"reflect"
	"testing"
)

func TestWallet(t *testing.T) {
	// Step 1: Initialize a new Wallet.
	w := New(100)

	// Step 2: Check initial balance.
	want := 100
	got := w.Balance()
	if got != want {
		t.Errorf("Balance() = %d, want %d", got, want)
	}

	// Step 3: Deposit.
	want = 150
	depositAmt := 50
	w.Deposit(depositAmt)
	got = w.Balance()
	if got != want {
		t.Errorf("Deposit() = %d, want %d", got, want)
	}

	// Step 4: Withdraw.
	want = 100
	withdrawAmt := 50
	w.Withdraw(withdrawAmt)
	got = w.Balance()
	if got != want {
		t.Errorf("Withdraw() = %d, want %d", got, want)
	}

	// Step 5: Get transactions.
	wantTransactions := []Transaction{
		{
			amount:    depositAmt,
			transType: Deposit,
		},
		{
			amount:    withdrawAmt,
			transType: Withdraw,
		},
	}
	gotTransactions := w.Transactions()
	if !reflect.DeepEqual(gotTransactions, wantTransactions) {
		t.Errorf("Transactions() = %v, \nwant %v", gotTransactions, wantTransactions)
	}
}
