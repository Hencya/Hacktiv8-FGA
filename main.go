package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	arr := []int{0x421, 0x51, 0x410, 0x51, 0x428, 0x51, 0x410, 0x51, 0x420, 0x51, 0x412, 0x51, 0x41e}

	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}

	for j := 0; j < 11; j++ {
		if j > 4 && j < 6 {
			for i, char := range arr {
				if i%2 == 0 {
					fmt.Printf("character %U %q starts at byte position %d\n", char, char, i)
				}
			}
			continue
		}
		fmt.Printf("Nilai j = %d\n", j)
	}
}
