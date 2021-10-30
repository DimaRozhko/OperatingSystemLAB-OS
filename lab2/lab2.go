package lab2

import (
	"fmt"
	"math/rand"
	"strconv"
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
var mailbox mailDTO = mailDTO{requestId: empty}

func userThreadA(moneyInAccount int) {
	sessionCounter++
	fmt.Println("THREAD A (start)")
	fmt.Printf("You balance in account:\t%d\n", moneyInAccount)
	fmt.Print("If you want to continue Yes/No [y/n]:")
	var getMoney string
	fmt.Scan(&getMoney)
	getMoney = strings.ToLower(getMoney)
	for getMoney == "y" || getMoney == "yes" {
		var getMoneyVal int
		fmt.Printf("How much money you want to get from ATM?\nInput:\t")
		fmt.Scan(&getMoneyVal)
		for mailbox.requestId != empty {
		}
		mailbox.transferMoney = getMoneyVal
		mailbox.moneyInAccount = moneyInAccount
		mailbox.requestId = get
		time.Sleep(1 * time.Microsecond)
		for mailbox.requestId != post {
		}
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Printf("|Mailbox after ATM transfer:\t\t\t\t\t\t\t|\n|->\t")
		fmt.Println(bankNoteInStock(mailbox.stock, "user") + "\t|")
		fmt.Printf("|->\tUSER GOT =========>%d$\t\t\t\t\t\t\t|\n", mailbox.transferMoney)
		fmt.Println("+-------------------------------------------------------------------------------+")
		moneyInAccount = mailbox.moneyInAccount
		mailbox.moneyInAccount = 0
		mailbox.transferMoney = 0
		mailbox.stock = stockMoney{}
		mailbox.requestId = empty
		// fmt.Printf("Set default mailbox:\n->\t")
		// bankNoteInStock(mailbox.stock, "mailbox")
		fmt.Printf("You balance in account:\t%d\n", moneyInAccount)
		fmt.Print("If you want to continue Yes/No [y/n]:")
		fmt.Scan(&getMoney)
	}
	sessionCounter--
	if sessionCounter == 0 {
		isATMwork = false
	}
	waittime.Done()
	fmt.Println("THREAD A (finish)")
}

func exchangeMoney(stock int, mailboxStock int, transferMoney int, bankNoteValue int) (int, int, int) {
	for transferMoney >= bankNoteValue && stock > 0 {
		transferMoney -= bankNoteValue
		stock--
		mailboxStock++
	}
	return stock, mailboxStock, transferMoney
}

func atmThreadB(stock stockMoney) {
	fmt.Println("THREAD B (start)")
	for isATMwork {
		if mailbox.requestId == get {
			mailbox.requestId = wait
			if mailbox.moneyInAccount >= mailbox.transferMoney {
				fmt.Printf("ATM before preparation:\n->\t")
				fmt.Println(bankNoteInStock(stock, "ATM"))
				var stockSave = stock
				var transferMoneyBuffer = mailbox.transferMoney
				stock.oneHundred, mailbox.stock.oneHundred, transferMoneyBuffer =
					exchangeMoney(stock.oneHundred, mailbox.stock.oneHundred, transferMoneyBuffer, 100)
				stock.fifty, mailbox.stock.fifty, transferMoneyBuffer =
					exchangeMoney(stock.fifty, mailbox.stock.fifty, transferMoneyBuffer, 50)
				stock.twenty, mailbox.stock.twenty, transferMoneyBuffer =
					exchangeMoney(stock.twenty, mailbox.stock.twenty, transferMoneyBuffer, 20)
				stock.ten, mailbox.stock.ten, transferMoneyBuffer =
					exchangeMoney(stock.ten, mailbox.stock.ten, transferMoneyBuffer, 10)
				stock.five, mailbox.stock.five, transferMoneyBuffer =
					exchangeMoney(stock.five, mailbox.stock.five, transferMoneyBuffer, 5)
				stock.two, mailbox.stock.two, transferMoneyBuffer =
					exchangeMoney(stock.two, mailbox.stock.two, transferMoneyBuffer, 2)
				stock.one, mailbox.stock.one, transferMoneyBuffer =
					exchangeMoney(stock.one, mailbox.stock.one, transferMoneyBuffer, 1)
				if transferMoneyBuffer != 0 {
					stock = stockSave
					mailbox.stock = stockMoney{one: 0, two: 0, five: 0, ten: 0, twenty: 0, fifty: 0, oneHundred: 0}
					mailbox.transferMoney = 0
					fmt.Println("Not enough money in the ATM:\t" + strconv.Itoa(transferMoneyBuffer))
				} else {
					mailbox.moneyInAccount -= mailbox.transferMoney
				}
				fmt.Printf("ATM after preparation:\n->\t")
				fmt.Println(bankNoteInStock(stock, "ATM"))
			} else {
				fmt.Printf("YOU CAN'T GET %d$:\n->\t", mailbox.transferMoney)
				fmt.Printf("In your account is only %d$\n", mailbox.moneyInAccount)
				mailbox.transferMoney = 0
			}
			mailbox.requestId = post
		}
	}
	waittime.Done()
	fmt.Println("THREAD B (finish)")
}

func bankNoteInStock(ATM stockMoney, stockName string) string {
	return fmt.Sprintf("Bank note in %s: 1$=%d, 2$=%d, 5$=%d, 10$=%d, 20$=%d, 50$=%d, 100$=%d",
		stockName, ATM.one, ATM.two, ATM.five, ATM.ten, ATM.twenty, ATM.fifty, ATM.oneHundred)
}
func Lab2() {
	fmt.Println("LAB2")
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

	fmt.Println(bankNoteInStock(ATM, "ATM"))

	isATMwork = true
	waittime.Add(2)
	go userThreadA(4000)
	go atmThreadB(ATM)

	waittime.Wait()
	fmt.Println("DONE")
}
