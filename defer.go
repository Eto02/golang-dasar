package main

import "fmt"

func simpleDefer() {
	defer fmt.Println("World") // Akan dieksekusi setelah fungsi selesai
	fmt.Println("Hello")
}

func main() {
	simpleDefer()
}
