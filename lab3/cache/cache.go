package cache

import (
	"fmt"
	"strconv"
	"strings"
)

var previousQuery string = ""
var previousQueryId int = 0
var kayTocacheDatMap []int
var cacheDatMap map[int][]string

func GetPreviousQuery() string {
	return previousQuery
}

func GetPreviousQueryId() int {
	return previousQueryId
}

func CheckIsCache(query string) bool {
	// fmt.Println(query)
	isCach := true
	if len(query) > 0 && strings.Compare(previousQuery, query) == 0 {
		return isCach
	}
	numQuery, _ := strconv.Atoi(query[1:])
	for _, num := range kayTocacheDatMap {
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
	return kayTocacheDatMap
}

func GetCacheDatMap() map[int][]string {
	return cacheDatMap
}

func CacheDatMapCreator(cacheDat []string) {
	cacheDatMap = make(map[int][]string)
	for _, tok := range cacheDat {
		if len(tok) < 5 {
			num, _ := strconv.ParseInt(tok, 2, 64)
			kayTocacheDatMap = append(kayTocacheDatMap, int(num))
		} else {
			cacheDatMap[kayTocacheDatMap[len(kayTocacheDatMap)-1]] =
				append(cacheDatMap[kayTocacheDatMap[len(kayTocacheDatMap)-1]], tok)
		}
	}
}
