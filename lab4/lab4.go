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
	s := strings.Split(string(buff), "\n")
	fmt.Println(s[0])
}
