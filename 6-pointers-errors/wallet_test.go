package pointersErrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		// Arrange
		wallet := Wallet{}

		// Act
		wallet.Deposit(Bitcoin(10))

		// Assert
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		// Arrange
		wallet := Wallet{balance: Bitcoin(25)}

		// Act
		err := wallet.Withdraw(Bitcoin(8))

		// Assert
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(17))
	})

	t.Run("Insufficient balance", func(t *testing.T) {
		// Arrange
		startingBalance := Bitcoin(25)
		wallet := Wallet{balance: Bitcoin(startingBalance)}

		// Act
		err := wallet.Withdraw(Bitcoin(9001))

		// Assert
		assertBalance(t, wallet, Bitcoin(startingBalance))
		assertError(t, err, ErrInufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("expected no error, got %q", got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error, got nil")
	}

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
