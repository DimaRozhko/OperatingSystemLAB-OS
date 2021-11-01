package lab3

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Println(buff)
	buff, _ = ioutil.ReadAll(mainmemDatFile)
	fmt.Println(buff)
	buff, _ = ioutil.ReadAll(cacheDatFile)
	fmt.Println(buff)

}
