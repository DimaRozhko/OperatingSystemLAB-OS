/**
Variant 13
*/

package main

import (
	"main/lab1"
	"main/lab2"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for _, arg := range argsWithoutProg {
		switch arg {
		case "lab1":
			lab1.Lab1()
		case "lab2":
			lab2.Lab2()
		}
	}
}
