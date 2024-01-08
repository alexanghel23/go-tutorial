package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

// Mutex = Mutable mutual exclusion, only allows one customer to acces the account at a time

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock() // lock the account so nobody else can access it
	defer a.lock.Unlock()

	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock() // unlock the account once the function is done

	if v > a.balance {
		pl("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn; Balance: %d\n", v, a.balance)
		a.balance -= v
	}

}

func main() {
	var acct Account
	acct.balance = 100

	pl("balance:", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}

	time.Sleep(2 * time.Second)
}
