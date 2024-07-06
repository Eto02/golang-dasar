package main

import "fmt"

func newSimpleDefer() {
	fmt.Println("Hello")
	message := recover()
	fmt.Println("Error", message)
}

func simplePanic() {
	defer newSimpleDefer()
	fmt.Println("Start")
	panic("Something went wrong!") // Memicu panic
	fmt.Println("End")             // Tidak akan dieksekusi karena panic
}

func main() {
	simplePanic()
	fmt.Println("Program continues after panic") // Tidak akan dieksekusi karena panic
}
