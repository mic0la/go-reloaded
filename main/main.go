package main

import (
	"fmt"
	"os"
	"reloaded"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid input")
		return
	}
	args := os.Args[1:]
	str, e := os.ReadFile("../texts/" + args[0])
	if e != nil {
		fmt.Println("invalid input")
		return
	}
	if len(str) == 0 {
		os.WriteFile("../texts/"+args[1], []byte{}, 0644)
		return
	}
	result := reloaded.CorrectAll(string(str))
	os.WriteFile("../texts/"+args[1], []byte(result), 0644)
}
