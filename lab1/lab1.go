// Process planning

package lab1

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type process struct {
	executTime int
	remainTime int
	name       string
	performed  bool
}

func processCreator(executTime int, remainTime int, name string) process {
	return process{
		executTime: executTime,
		remainTime: remainTime,
		name:       name,
		performed:  false,
	}
}

var queue1RR, queue2RR, queue3SRTF []process

var waittime sync.WaitGroup

var mutex sync.Mutex

const (
	numQueueProcess = 3
	quantQueue1     = 50
	quantQueue2     = 100
	maxExecutTime   = 200
	minRemainTime   = 1000
	maxRemainTime   = 5000
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

// func runProcess(queue []process, index int) {
// 	queue[index] = processCreator(queue[index].executTime,
// 		queue[index].remainTime-queue[index].executTime, queue[index].name)
// }

func queueThreadRR(queue []process) {
	var emptyProcessCounter int = 0
	for emptyProcessCounter != len(queue) {
		emptyProcessCounter = 0
		for i, process := range queue {
			if process.remainTime > 0 {
				time.Sleep(time.Duration(process.executTime) * time.Millisecond)
				fmt.Print(process.name + "\t")
				fmt.Print(process)
				fmt.Print("\t->\t")
				mutex.Lock()
				process.remainTime = process.remainTime - process.executTime
				queue[i].performed = true
				queue[i].remainTime = process.remainTime
				queue[i].performed = false
				mutex.Unlock()
				fmt.Println(process)
				// fmt.Print("\tQ3")
				// fmt.Println(queue3SRTF)
			} else {
				emptyProcessCounter++
			}
		}
	}
	for i, process := range queue3SRTF {
		if process.executTime > 0 && !process.performed {
			queue3SRTF[i] = queue[0]
			queue[0] = process
			queue[0].executTime = queue[1].executTime
			for queue[0].remainTime > 0 {
				time.Sleep(time.Duration(time.Duration(queue[0].executTime) * time.Millisecond))
				fmt.Print(queue[0].name + "\t")
				fmt.Print(queue[0])
				mutex.Lock()
				queue[0].remainTime = queue[0].remainTime - queue[0].executTime
				mutex.Unlock()
				fmt.Print("\t->\t")
				fmt.Println(queue[0])
			}
		}
	}
	waittime.Done()
}

func queueThreadSRTF(queue []process) {
	var (
		executionIdQ3SRTF   = 0
		minRemainQ3SRTF     = 0
		emptyProcessCounter = 0
	)
	for emptyProcessCounter != len(queue) {
		minRemainQ3SRTF = maxRemainTime
		for i, process := range queue {
			if process.remainTime < minRemainQ3SRTF && process.remainTime > 0 {
				executionIdQ3SRTF = i
				minRemainQ3SRTF = process.remainTime
			}
		}
		queue[executionIdQ3SRTF].performed = true
		for queue[executionIdQ3SRTF].remainTime > 0 {
			time.Sleep(time.Duration(queue[executionIdQ3SRTF].executTime) * time.Millisecond)
			fmt.Print(queue[executionIdQ3SRTF].name + "\t")
			fmt.Print(queue[executionIdQ3SRTF])
			mutex.Lock()
			queue[executionIdQ3SRTF].remainTime = queue[executionIdQ3SRTF].remainTime - queue[executionIdQ3SRTF].executTime
			mutex.Unlock()
			fmt.Print("\t->\t")
			fmt.Println(queue[executionIdQ3SRTF])
			// fmt.Print("\tQ1")
			// fmt.Println(queue1RR)
			// fmt.Print("\tQ2")
			// fmt.Println(queue2RR)
		}
		queue[executionIdQ3SRTF].performed = false
		emptyProcessCounter++
	}
	waittime.Done()
}

func Lab1() {
	fmt.Println("LAB1")

	rand.Seed(time.Now().UnixNano())
	fmt.Print("Q3")
	fmt.Println(queue3SRTF)
	for i := 0; i < numQueueProcess; i++ {
		queue1RR = append(queue1RR,
			processCreator(quantQueue1, rand.Intn(maxRemainTime-minRemainTime)+minRemainTime, nameQ1RR+strconv.Itoa(i)))
		queue2RR = append(queue2RR,
			processCreator(quantQueue2, rand.Intn(maxRemainTime-minRemainTime)+minRemainTime, nameQ2RR+strconv.Itoa(i)))
		queue3SRTF = append(queue3SRTF,
			processCreator(rand.Intn(maxExecutTime-quantQueue1)+quantQueue1,
				rand.Intn(maxRemainTime-minRemainTime)+minRemainTime, nameQ3SRTF+strconv.Itoa(i)))
	}

	printAllQueue(queue1RR, queue2RR, queue3SRTF)

	waittime.Add(3)

	go queueThreadRR(queue1RR)
	go queueThreadRR(queue2RR)
	go queueThreadSRTF(queue3SRTF)

	waittime.Wait()

	fmt.Println("done")
	printAllQueue(queue1RR, queue2RR, queue3SRTF)
}
