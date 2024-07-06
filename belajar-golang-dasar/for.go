package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	// Iterasi melalui array atau slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterasi melalui string
	str := "hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// Iterasi melalui map
	myMap := map[string]int{"a": 1, "b": 2}
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	neCounter := 1
	for {
		fmt.Println("Infinite Loop")
		// Tambahkan kondisi untuk keluar dari loop jika diperlukan
		if neCounter == 3 {
			break
		}
		fmt.Println("Perulangan ke", neCounter)
		neCounter++

	}
}
