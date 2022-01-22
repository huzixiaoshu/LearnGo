package main

import (
	//"fmt"
	"testing"
)

// func TestWallet(t *testing.T) {

// 	t.Run("Deposit", func(t *testing.T) {
// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()

// 		//fmt.Println("address of balance in test is", &wallet.balance)

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}

// 		wallet.Withdraw(10)

// 		got := wallet.Balance()

// 		//fmt.Println("address of balance in test is", &wallet.balance)

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// }

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("didn't get an error but wanted one")
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")

	})
}
