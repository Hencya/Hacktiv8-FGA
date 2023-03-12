package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	inputText := "selamat malam"
	mapText := make(map[string]int)
	for _, c := range inputText {
		fmt.Printf("%v\n", string(c))
		mapText[string(c)]++
	}
	fmt.Println(mapText)
}
