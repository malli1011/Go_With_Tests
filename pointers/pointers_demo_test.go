package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		if err != nil {
			t.Fatal("Expect no error to be thrown")
		}

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(30))

		asserError(t, err, "canot withdraw, insufficient funds")
		assertBalance(t, wallet, Bitcoin(20))
	})

}

func asserError(t testing.TB, got error, expected string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but expected one")
	}
	if got.Error() != expected {
		t.Errorf("got %q, but expected %q", got, expected)
	}
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != expected {
		t.Errorf("got %s but expected %s", got, expected)
	}
}
