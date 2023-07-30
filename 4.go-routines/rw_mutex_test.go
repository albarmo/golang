package go_routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.balance = account.balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBallance() int {
	account.RWMutex.RLock()
	balance := account.balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			account.AddBalance(1)
			fmt.Println(account.GetBallance())
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Ballance : ", account.GetBallance())
}
