package wallet

import (
	"reflect"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		w := New(100)

		want := 100
		got := w.Balance()
		if got != want {
			t.Errorf("Balance() = %d, want %d", got, want)
		}
	})

	t.Run("Balance", func(t *testing.T) {
		w := New(100)
		want := 150
		w.Deposit(50)
		got := w.Balance()
		if got != want {
			t.Errorf("Deposit() = %d, want %d", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := New(100)
		want := 50
		w.Withdraw(50)
		got := w.Balance()
		if got != want {
			t.Errorf("Withdraw() = %d, want %d", got, want)
		}
	})

	t.Run("Transactions", func(t *testing.T) {
		w := New(100)
		depositAmt := 50
		w.Deposit(depositAmt)
		withdrawAmt := 50
		w.Withdraw(withdrawAmt)

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
	})
}
