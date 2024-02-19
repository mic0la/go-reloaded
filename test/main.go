package main

import (
	"cadet"
	"fmt"
	"os"
)

func main() {
	//s := "4B (hex) points get"
	s, err := os.ReadFile("../cat.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("test value:", string(cadet.HexHandler(s)))
}
