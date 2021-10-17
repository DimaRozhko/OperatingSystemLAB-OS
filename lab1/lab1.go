// Process planning

package lab1

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type process struct {
	executTime int
	remainTime int
	name       string
}

func processCreator(executTime int, remainTime int, name string) process {
	return process{
		executTime: executTime,
		remainTime: remainTime,
		name:       name,
	}
}

var queue1RR, queue2RR, queue3SRTF []process

const (
	numQueueProcess = 3
	minExecutTime   = 50
	maxExecutTime   = 1000
	minRemainTime   = 1000
	maxRemainTime   = 1000000
	nameQ1RR        = "Q1RR"
	nameQ2RR        = "Q2RR"
	nameQ3SRTF      = "Q3SRTF"
	rowNameCell     = "+-------+"
	rowTableCell    = "+---------------------+"
)

func printRowTable(queue []process, name string) {
	fmt.Print(rowNameCell)
	for i := 0; i < numQueueProcess; i++ {
		fmt.Print(rowTableCell)
	}
	fmt.Print("\n|" + name + "\t|")
	var nameFormat string = "|"
	for _, process := range queue {
		if len(process.name) < 7 {
			nameFormat = "  |"
		} else {
			nameFormat = "|"
		}
		fmt.Print("|" + strconv.Itoa(process.executTime) + "   |" +
			strconv.Itoa(process.remainTime) + "|" + process.name + nameFormat)
	}
	fmt.Println()
	fmt.Print(rowNameCell)
	for i := 0; i < numQueueProcess; i++ {
		fmt.Print(rowTableCell)
	}
	fmt.Println()
}

func printAllQueue(queue1 []process, queue2 []process, queue3 []process) {
	fmt.Print(rowNameCell)
	for i := 0; i < numQueueProcess; i++ {
		fmt.Print(rowTableCell)
	}
	fmt.Print("\n|QUEU NM|")
	for i := 0; i < numQueueProcess; i++ {
		fmt.Print("|EXE TM|REM TM|P NAME |")
	}
	fmt.Println()
	printRowTable(queue1, nameQ1RR)
	printRowTable(queue2, nameQ2RR)
	printRowTable(queue3, nameQ3SRTF)
}

func Lab1() {
	fmt.Println("LAB1")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numQueueProcess; i++ {
		queue1RR = append(queue1RR,
			processCreator(rand.Intn(maxExecutTime-minExecutTime)+minExecutTime,
				rand.Intn(maxRemainTime-minRemainTime)+minRemainTime, nameQ1RR+strconv.Itoa(i)))
		queue2RR = append(queue2RR,
			processCreator(rand.Intn(maxExecutTime-minExecutTime)+minExecutTime,
				rand.Intn(maxRemainTime-minRemainTime)+minRemainTime, nameQ2RR+strconv.Itoa(i)))
		queue3SRTF = append(queue3SRTF,
			processCreator(rand.Intn(maxExecutTime-minExecutTime)+minExecutTime,
				rand.Intn(maxRemainTime-minRemainTime)+minRemainTime, nameQ3SRTF+strconv.Itoa(i)))
	}
	printAllQueue(queue1RR, queue2RR, queue3SRTF)
}
