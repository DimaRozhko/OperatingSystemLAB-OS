package mainmem

type memRow struct {
	num   string
	value string
}

var mainDatMap []memRow

func GetMainDatMap() []memRow {
	return mainDatMap
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
