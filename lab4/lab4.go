package lab4

import (
	"fmt"
	"io/ioutil"
	fat16 "main/lab4/FAT16"
	"os"
	"strings"
)

func Lab4() {
	fmt.Println("LAB4")
	file, err := os.OpenFile("lab4/data/fileinfo.dat", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	buff, _ := ioutil.ReadAll(file)
	fat16.TableCreator()
	splitRow := strings.Split(string(buff), "\n")
	for i, row := range splitRow {
		if i > 0 {
			break
		}
		fileInfo := strings.Split(row, "$")
		fat16.AddFileToFAT16Table(fileInfo[0], fileInfo[1], fileInfo[2], fileInfo[3], fileInfo[4])
	}
	file.Close()
}
