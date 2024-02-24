package main

import (
	"fmt"
	"os"
	"reloaded"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	primaryStr, _ := os.ReadFile("../texts/" + args[0])
	secondaryStr := reloaded.CorrectAll(string(primaryStr))
	//fmt.Println(secondaryStr)
	err := os.WriteFile("../texts/"+args[1], []byte(secondaryStr), 0644)
	check(err)
	fmt.Println()
}
