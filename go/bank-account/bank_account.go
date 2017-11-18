package account

import (
	"sync"
)

type Account struct {
	sync.Mutex
	balance int64
	open    bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, open: true}
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return a.balance, false
	}

	payout := a.balance
	a.balance = 0
	a.open = false

	return payout, true
}

func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return a.balance, false
	}

	newBalance := a.balance + amount
	if newBalance < 0 {
		return a.balance, false
	}

	a.balance = newBalance

	return a.balance, true
}
