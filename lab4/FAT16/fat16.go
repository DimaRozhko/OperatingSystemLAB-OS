package fat16

import (
	"fmt"
)

type clusterCode struct {
	word byte
}

const (
	clusterLength      = 0xFF + 1
	usageClusterLength = 0x0F
	firstUsageCluster  = 0x06 // only for first row in the table
)

var fat16Table [clusterLength][clusterLength]clusterCode

// var fat16ReserveTable [clusterLength][clusterLength]clusterCode

func TableCreator() {
	// Set necessary file to run OS
	fat16Table[0x00][0x01] = clusterCode{word: 0xff}
	fat16Table[0x00][0x02] = clusterCode{word: 0x03}
	fat16Table[0x00][0x03] = clusterCode{word: 0x04}
	fat16Table[0x00][0x04] = clusterCode{word: 0x05}
	fat16Table[0x00][0x05] = clusterCode{word: 0xff}
	fmt.Println(fat16Table[0])
}

func AddFileToFAT16Table() {

}
