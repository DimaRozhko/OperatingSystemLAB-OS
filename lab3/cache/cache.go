package cache

import (
	"fmt"
	"strconv"
	"strings"
)

var previousQuery string = ""
var kayTocacheDatMap []int
var cacheDatMap map[int][]string

func GetPreviousQuery() string {
	return previousQuery
}

func CheckIsCache(query string) bool {
	// fmt.Println(query)
	if len(query) > 0 && strings.Compare(previousQuery, query) == 0 {
		return true
	}
	numQuery, _ := strconv.Atoi(query[1:])
	binNumQuery := strconv.FormatInt(int64(numQuery), 2)
	fmt.Println(strings.Repeat("0", 24-len(binNumQuery)) + binNumQuery)
	previousQuery = query
	return false
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
