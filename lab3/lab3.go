package lab3

import (
	"fmt"
	"io/ioutil"
	"main/lab3/cache"
	"main/lab3/mainmem"
	"os"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

func openFileWithCheck(name string, flag int) *os.File {
	file, err := os.OpenFile(name, flag, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	return file
}

func moveDataFromCacheToMainMemory() {
	mainmem.AddDataFromCache(cache.GetCacheDatMap(), cache.GetKayTocacheDatMap())
	cache.ClearСache()
}

func queryRunner(originData []string) {
	for _, tok := range originData {
		if cache.IsCacheFull() {
			moveDataFromCacheToMainMemory()
		}
		if govalidator.IsInt(tok) {
			for i, _ := strconv.ParseInt(tok, 10, 64); i > 0; i-- {
				if cache.IsCacheFull() {
					moveDataFromCacheToMainMemory()
				}
				cache.AddRandomDataCacheDatMapByKey(cache.GetPreviousQueryId())
				fmt.Println("CACHE " + cache.GetPreviousQuery())
				fmt.Println(cache.GetCacheDatMap())

			}
		} else {
			if cache.CheckIsCache(tok) {
				fmt.Println("CACHE " + cache.GetPreviousQuery())
				cache.AddRandomDataCacheDatMapByKey(cache.GetPreviousQueryId())
				fmt.Println(cache.GetCacheDatMap())
			} else {
				fmt.Println("NOT CACHE " + tok)
				mainmem.AddRandomToMainMem(tok[1:])
				cache.AddKeyToCacheMap(tok[1:])
			}
		}
	}
	moveDataFromCacheToMainMemory()
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
	buff, _ = ioutil.ReadAll(mainmemDatFile)
	mainmemDat := strings.Fields(string(buff))
	mainmem.MainMemDatMapCreator(mainmemDat)
	fmt.Println("INPUT MEMORY CONDIRION")
	fmt.Println(mainmem.GetMainDatMap())
	buff, _ = ioutil.ReadAll(cacheDatFile)

	cache.CacheDatMapCreator(strings.Fields(string(buff)))
	fmt.Println("INPUT CACHE CONDIRION")
	fmt.Println(cache.GetCacheDatMap())

	queryRunner(originData)

	fmt.Println(mainmem.GetMainDatMap())

}
