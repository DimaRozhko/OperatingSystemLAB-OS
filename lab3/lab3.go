package lab3

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func openFileWithCheck(name string, flag int) *os.File {
	file, err := os.OpenFile(name, flag, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	return file
}

func Lab3() {
	fmt.Println("LAB3")
	os.MkdirAll("lab3/log", os.ModePerm)

	// cacheLogFile := openFileWithCheck("lab3/log/cache.log", os.O_CREATE|os.O_WRONLY)
	// mainmemLogFile := openFileWithCheck("lab3/log/mainmem.log", os.O_CREATE|os.O_WRONLY)
	originDatFile := openFileWithCheck("lab3/data/origin.dat", os.O_RDONLY)
	mainmemDatFile := openFileWithCheck("lab3/data/mainmem.dat", os.O_RDONLY)
	cacheDatFile := openFileWithCheck("lab3/data/cache.dat", os.O_RDONLY)

	buff, _ := ioutil.ReadAll(originDatFile)
	originData := strings.Fields(string(buff))
	fmt.Println(originData)
	buff, _ = ioutil.ReadAll(mainmemDatFile)
	mainmemDat := strings.Fields(string(buff))
	fmt.Println(mainmemDat)
	var kayToMainMemMap []int
	mainMemMap := make(map[int]string)
	for i, tok := range mainmemDat {
		if i%2 == 0 {
			num, _ := strconv.ParseInt(tok, 2, 64)
			kayToMainMemMap = append(kayToMainMemMap, int(num))
		} else {
			mainMemMap[kayToMainMemMap[len(kayToMainMemMap)-1]] = tok
		}
	}
	fmt.Println(kayToMainMemMap)
	fmt.Println(mainMemMap)
	buff, _ = ioutil.ReadAll(cacheDatFile)
	cacheDat := strings.Fields(string(buff))
	fmt.Println(cacheDat)

}
