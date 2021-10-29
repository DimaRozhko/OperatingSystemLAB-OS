package lab2

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type stockMoney struct {
	one        int
	two        int
	five       int
	ten        int
	twenty     int
	fifty      int
	oneHundred int
}

const (
	get   = iota
	post  = iota
	wait  = iota
	empty = iota
)

type mailDTO struct {
	stock          stockMoney
	moneyInAccount int
	transferMoney  int
	requestId      int
}

var sessionCounter = 0
var isATMwork = false

var waittime sync.WaitGroup

var mailbox mailDTO = mailDTO{
	requestId: empty,
}

func userThreadA(moneyInAccount int) {
	sessionCounter++
	fmt.Println("THREAD A")
	fmt.Print("If you want to continue Yes/No [y/n]:")
	var getMoney string
	fmt.Scan(&getMoney)
	getMoney = strings.ToLower(getMoney)
	for getMoney == "y" || getMoney == "yes" {
		fmt.Printf("You balance in account:\t%d\nHow much money you want to get from ATM?\nInput:\t", moneyInAccount)
		var getMoneyVal int
		fmt.Scan(&getMoneyVal)
		fmt.Println(getMoneyVal)
		for mailbox.requestId != empty {
		}
		// if mailbox.requestId == empty {
		mailbox.transferMoney = getMoneyVal
		mailbox.moneyInAccount = moneyInAccount
		mailbox.requestId = get
		fmt.Println(mailbox)
		for mailbox.requestId != post {
		}
		fmt.Println(mailbox)
		moneyInAccount = mailbox.moneyInAccount
		mailbox.moneyInAccount = 0
		mailbox.transferMoney = 0
		mailbox.stock = stockMoney{}
		mailbox.requestId = empty
		fmt.Println(mailbox)
		fmt.Print("If you want to continue Yes/No [y/n]:")
		fmt.Scan(&getMoney)
	}
	sessionCounter--
	if sessionCounter == 0 {
		isATMwork = false
	}
	waittime.Done()
}

func atmThreadB(stock stockMoney) {
	fmt.Println("THREAD B")
	for isATMwork {
		if mailbox.requestId == get {
			mailbox.requestId = wait

			mailbox.requestId = post
		}
	}
	waittime.Done()
}

func Lab2() {
	rand.Seed(time.Now().UnixNano())
	var capacityATM = 1000
	var ATM stockMoney = stockMoney{
		one:        rand.Intn(capacityATM),
		two:        rand.Intn(capacityATM),
		five:       rand.Intn(capacityATM),
		ten:        rand.Intn(capacityATM),
		twenty:     rand.Intn(capacityATM),
		fifty:      rand.Intn(capacityATM),
		oneHundred: rand.Intn(capacityATM),
	}

	fmt.Println(ATM)

	isATMwork = true
	waittime.Add(2)
	go userThreadA(4000)
	go atmThreadB(ATM)

	waittime.Wait()
	fmt.Println("DONE")
}
