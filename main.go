/**
Variant 13
*/

package main

import (
	"main/lab1"
	"main/lab2"
	"main/lab3"
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
		case "lab3":
			lab3.Lab3()
		case "lab4":
			lab3.Lab3()
		}
	}
}
