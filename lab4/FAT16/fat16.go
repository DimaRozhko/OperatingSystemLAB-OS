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

type clusterStorage struct {
	id       int
	fileNeme string
}

const (
	clusterLength          = 0xFF + 1
	usedClusterLength      = 0x0F
	firstUsedCluster       = 0x06 // only for first row in the table
	defectClusterCode      = 0xF7
	percentOfDefectCluster = 51
	sizeCluster            = 2 // 2Kb

)

var (
	fileIds                 []int
	curentFileName          string
	fileId                  int
	clustereMap             map[clusterCode]clusterStorage
	fat16Table              [clusterLength][clusterLength]clusterCode
	currentRowFAT16Table    byte = 0x00
	currentColumnFAT16Table byte = 0x00
)

func GetAllFileId() []int {
	return fileIds
}

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

func PrintCluster() {
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
	fmt.Println("INIT FAT16 TABLE -> Set necessary file to run OS:")
	fmt.Print("First line in \"used\" cluster:\t")
	fmt.Println(fat16Table[currentRowFAT16Table][:usedClusterLength])
	setDefectCluster()
	PrintCluster()
	clustereMap = make(map[clusterCode]clusterStorage)
}

func setCluster() {
	var (
		rowToUseNextCluster    byte = currentRowFAT16Table
		columnToUseNextCluster byte = currentColumnFAT16Table + 1
		facedDefect            bool = false
	)
	for fat16Table[rowToUseNextCluster][columnToUseNextCluster].word[0] == 0xFF &&
		fat16Table[rowToUseNextCluster][columnToUseNextCluster].word[1] == defectClusterCode &&
		columnToUseNextCluster < usedClusterLength && rowToUseNextCluster < usedClusterLength {
		fmt.Println(fat16Table[rowToUseNextCluster][columnToUseNextCluster])

		if columnToUseNextCluster == usedClusterLength-1 {
			rowToUseNextCluster++
			columnToUseNextCluster = 0
		} else {
			columnToUseNextCluster++
		}
		facedDefect = true
	}
	if facedDefect {
		if columnToUseNextCluster == usedClusterLength-1 {
			rowToUseNextCluster++
			columnToUseNextCluster = 0
		} else {
			columnToUseNextCluster++
		}
	}
	fat16Table[currentRowFAT16Table][currentColumnFAT16Table].word = [2]byte{rowToUseNextCluster, columnToUseNextCluster}
	clustereMap[clusterCode{word: [2]byte{rowToUseNextCluster, columnToUseNextCluster}}] = clusterStorage{id: fileId, fileNeme: curentFileName}
	// currentRowFAT16Table = rowToUseNextCluster
	// currentColumnFAT16Table = columnToUseNextCluster
}

func moveFileToFAT16Table(fileSizeDec int) {

	currentColumnFAT16Table++
	// if currentColumnFAT16Table == usedClusterLength-1 {
	// 	fat16Table[currentRowFAT16Table][currentColumnFAT16Table].word = [2]byte{currentRowFAT16Table + 1, 0}
	// 	currentColumnFAT16Table = 0
	// 	currentRowFAT16Table++
	// }
	for ; currentRowFAT16Table < usedClusterLength; currentRowFAT16Table++ {
		for ; currentColumnFAT16Table < usedClusterLength; currentColumnFAT16Table++ {
			if fat16Table[currentRowFAT16Table][currentColumnFAT16Table].word[0] == 0xFF &&
				(fat16Table[currentRowFAT16Table][currentColumnFAT16Table].word[1] == defectClusterCode) {

				if currentColumnFAT16Table == usedClusterLength-1 {
					currentRowFAT16Table++
				}
				continue
			}
			if currentColumnFAT16Table == usedClusterLength-1 {
				fat16Table[currentRowFAT16Table][currentColumnFAT16Table].word = [2]byte{currentRowFAT16Table + 1, 0}
				clustereMap[clusterCode{word: [2]byte{currentRowFAT16Table + 1, 0}}] = clusterStorage{id: fileId, fileNeme: curentFileName}
				currentColumnFAT16Table = 0
				currentRowFAT16Table++
			}
			if fileSizeDec >= sizeCluster {
				setCluster()

			} else {
				if fileSizeDec > 0 {
					setCluster()
				} else {
					return
				}
			}
			fileSizeDec -= sizeCluster
		}

	}
}

func CreateFileInFAT16Table(fileName string, attribute string, creationTime string, creationDate string, fileSizeHex string) {
	fileId = len(fileIds)
	curentFileName = fileName
	fmt.Println("Fileneme: " + fileName)
	fmt.Println("\tAttribute: " + attribute)
	fmt.Println("\tCreation date: " + creationDate)
	fmt.Println("\tCreation time: " + creationTime)
	// fmt.Println("\tCluster number: " + clusterNumHex)
	fmt.Println("\tFile size: " + fileSizeHex)
	fileSizeDec, _ := strconv.ParseInt(fileSizeHex, 16, 64)
	moveFileToFAT16Table(int(fileSizeDec))
	fat16Table[currentRowFAT16Table][currentColumnFAT16Table] = clusterCode{word: [2]byte{0xFF, 0xff}}
	PrintCluster()
	fileIds = append(fileIds, fileId)
}

func FindFileById(id int) {
	var codeHex string
	isFind := false
	for i := 0; i < usedClusterLength; i++ {
		for j := 0; j < usedClusterLength; j++ {
			if value, ok := clustereMap[fat16Table[i][j]]; ok {
				if value.id == id {
					isFind = true
					fmt.Print("FOUND: ")
					fmt.Print("Cluster: ")
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
					fmt.Println("\t->\tFilaname: ", value.fileNeme)
				}
			}
		}
	}
	if !isFind {
		fmt.Printf("CANNOT MOUNF FILE BY ID = %d\n", id)
	}
}

func DeleteFileById(id int) {
	var codeHex string
	for i := 0; i < usedClusterLength; i++ {
		for j := 0; j < usedClusterLength; j++ {
			if value, ok := clustereMap[fat16Table[i][j]]; ok {
				if value.id == id {
					fmt.Print("DELETED: ")
					fmt.Print("Cluster: ")
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
					fmt.Println("\t->\tFilaname: ", value.fileNeme)
					delete(clustereMap, fat16Table[i][j])
					fat16Table[i][j] = clusterCode{word: [2]byte{0x00, 0x00}}
				}
			}
		}
	}
}
