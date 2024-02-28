package main

import (
	"os"
	"reloaded"
)

func main() {
	args := os.Args[1:]
	str, _ := os.ReadFile("../texts/" + args[0])
	result := reloaded.CorrectAll(string(str))
	os.WriteFile("../texts/"+args[1], []byte(result), 0644)
}
