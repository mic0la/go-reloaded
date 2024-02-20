package main

import (
	"cadet"
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
	secondaryStr := cadet.CorrectAll(string(primaryStr))
	err := os.WriteFile("../texts/"+args[1], []byte(secondaryStr), 0644)
	check(err)
}
