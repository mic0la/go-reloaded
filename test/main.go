package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	primaryStr, _ := os.ReadFile("../texts/" + args[0])
	err := os.WriteFile("../texts/"+args[1], primaryStr, 0644)
	check(err)
}
