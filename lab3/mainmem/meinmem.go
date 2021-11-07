package mainmem

import (
	"math/rand"
	"strconv"
	"time"
)

type memRow struct {
	num   string
	value string
}

var mainDatMap []memRow

func GetMainDatMap() []memRow {
	return mainDatMap
}

func AddRandomToMainMem(num string) {
	rand.Seed(time.Now().UnixNano())
	convNum, _ := strconv.ParseInt(num, 10, 64)
	mainDatMap = append(mainDatMap, memRow{num: strconv.FormatInt(convNum, 2)[:4], value: strconv.FormatInt(rand.Int63(), 2)[:24]})
}

func MainMemDatMapCreator(mainDat []string) {
	for _, tok := range mainDat {
		if len(tok) < 5 {
			mainDatMap = append(mainDatMap, memRow{num: tok})
		} else {
			mainDatMap[len(mainDatMap)-1].value = tok
		}
	}
}
