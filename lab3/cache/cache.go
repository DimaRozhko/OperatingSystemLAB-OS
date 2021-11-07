package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var previousQuery string = ""
var previousQueryId int = 0
var kayToCacheDatMap []int
var cacheDatMap map[int][]string
var cacheIsModify bool
var cacheCounter int = 0

func GetCacheCounter() int {
	return cacheCounter
}

func IsCacheFull() bool {
	if cacheCounter < 10 {
		return false
	} else {
		return true
	}
}

func GetPreviousQuery() string {
	return previousQuery
}

func GetPreviousQueryId() int {
	return previousQueryId
}

func GetCacheIsModify() bool {
	return cacheIsModify
}

func AddKeyToCacheMap(keyNum string) {
	convNum, _ := strconv.ParseInt(keyNum, 10, 64)
	kayToCacheDatMap = append(kayToCacheDatMap, int(convNum))
}

func AddRandomDataCacheDatMapByKey(key int) {
	rand.Seed(time.Now().UnixNano())
	cacheDatMap[key] = append(cacheDatMap[key], strconv.FormatInt(rand.Int63(), 2)[:24])
	cacheIsModify = true
	cacheCounter++
}

func CheckIsCache(query string) bool {
	isCach := true
	if len(query) > 0 && strings.Compare(previousQuery, query) == 0 {
		return isCach
	}
	numQuery, _ := strconv.Atoi(query[1:])
	for _, num := range kayToCacheDatMap {
		if num == numQuery {
			isCach = false
			break
		}
	}
	if isCach {
		return !isCach
	}
	isCach = true
	binNumQuery := strconv.FormatInt(int64(numQuery), 2)
	fmt.Println(strings.Repeat("0", 4-len(binNumQuery)) + binNumQuery)
	previousQuery = query
	previousQueryId = numQuery
	return isCach
}

func GetKayTocacheDatMap() []int {
	return kayToCacheDatMap
}

func GetCacheDatMap() map[int][]string {
	return cacheDatMap
}

func CacheDatMapCreator(cacheDat []string) {
	cacheDatMap = make(map[int][]string)
	for _, tok := range cacheDat {
		if len(tok) < 5 {
			num, _ := strconv.ParseInt(tok, 2, 64)
			kayToCacheDatMap = append(kayToCacheDatMap, int(num))
		} else {
			cacheIsModify = true
			cacheDatMap[kayToCacheDatMap[len(kayToCacheDatMap)-1]] =
				append(cacheDatMap[kayToCacheDatMap[len(kayToCacheDatMap)-1]], tok)
		}
	}
}

func ClearСache() {
	for _, numKey := range kayToCacheDatMap {
		if len(cacheDatMap[numKey]) > 0 {
			cacheDatMap[numKey] = make([]string, 0)
		}
	}
	cacheCounter = 0
}
