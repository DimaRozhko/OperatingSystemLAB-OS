package fat16

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type clusterCode struct {
	word [2]byte
}

const (
	clusterLength          = 0xFF + 1
	usedClusterLength      = 0x0F
	firstUsedCluster       = 0x06 // only for first row in the table
	defectClusterCode      = 0xF7
	percentOfDefectCluster = 0
	sizeCluster            = 2 // 2Kb

)

var (
	fat16Table              [clusterLength][clusterLength]clusterCode
	currentRowFAT16Table    = 0x00
	currentColumnFAT16Table = 0x00
)

// var fat16ReserveTable [clusterLength][clusterLength]clusterCode

func setDefectCluster() {
	rand.Seed(time.Now().UnixNano())
	for i := 0x01; i < usedClusterLength; i++ {
		for j := 0x00; j < usedClusterLength; j++ {
			if rand.Intn(percentOfDefectCluster+1)+0 > 100-percentOfDefectCluster {
				fat16Table[i][j] = clusterCode{word: [2]byte{0xFF, defectClusterCode}}
			}
		}
	}
}

func printCluster() {
	fmt.Println("Cluster status:")
	var codeHex string
	for i := 0x00; i < usedClusterLength; i++ {
		for j := 0x00; j < usedClusterLength; j++ {
			fmt.Print(" ")
			codeHex = strconv.FormatInt(int64(fat16Table[i][j].word[0]), 16)
			if len(codeHex) == 1 {
				fmt.Print("0")
			}
			fmt.Print(strings.ToUpper(codeHex))
			codeHex = strconv.FormatInt(int64(fat16Table[i][j].word[1]), 16)
			if len(codeHex) == 1 {
				fmt.Print("0")
			}
			fmt.Print(strings.ToUpper(codeHex))
		}
		fmt.Println()
	}
}

func TableCreator() {
	// Set necessary file to run OS
	currentColumnFAT16Table++
	fat16Table[currentRowFAT16Table][currentColumnFAT16Table] = clusterCode{word: [2]byte{0x00, 0xff}}
	currentColumnFAT16Table++
	fat16Table[currentRowFAT16Table][currentColumnFAT16Table] = clusterCode{word: [2]byte{0x00, 0x03}}
	currentColumnFAT16Table++
	fat16Table[currentRowFAT16Table][currentColumnFAT16Table] = clusterCode{word: [2]byte{0x00, 0x04}}
	currentColumnFAT16Table++
	fat16Table[currentRowFAT16Table][currentColumnFAT16Table] = clusterCode{word: [2]byte{0x00, 0x05}}
	currentColumnFAT16Table++
	fat16Table[currentRowFAT16Table][currentColumnFAT16Table] = clusterCode{word: [2]byte{0xFF, 0xff}}
	currentColumnFAT16Table++
	fmt.Println("INIT FAT16 TABLE -> Set necessary file to run OS:")
	fmt.Print("First line in \"used\" cluster:\t")
	fmt.Println(fat16Table[currentRowFAT16Table][:usedClusterLength])
	setDefectCluster()
	printCluster()
}

func setCluster() {

}

func moveFileToFAT16Table(fileSizeDec int) {
	for i := currentRowFAT16Table; i < usedClusterLength; i++ {
		for j := currentColumnFAT16Table; j < usedClusterLength; j++ {
			if fileSizeDec >= sizeCluster {
				setCluster()
			} else {
				if fileSizeDec > 0 {
					setCluster()
				}
				if j == usedClusterLength-1 {
					currentRowFAT16Table = i + 1
					currentColumnFAT16Table = 0
				} else {
					currentRowFAT16Table = i
					currentColumnFAT16Table = j
				}
				return
			}
			fileSizeDec -= sizeCluster
		}
	}

}

func AddFileToFAT16Table(fileName string, attribute string, creationTime string, creationDate string, fileSizeHex string) {
	fmt.Println("Fileneme: " + fileName)
	fmt.Println("\tAttribute: " + attribute)
	fmt.Println("\tCreation date: " + creationDate)
	fmt.Println("\tCreation time: " + creationTime)
	// fmt.Println("\tCluster number: " + clusterNumHex)
	fmt.Println("\tFile size: " + fileSizeHex)
	fileSizeDec, _ := strconv.ParseInt(fileSizeHex, 16, 64)
	moveFileToFAT16Table(int(fileSizeDec))

}
