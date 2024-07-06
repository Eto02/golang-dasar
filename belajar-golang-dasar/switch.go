package main

import "fmt"

func main() {
	name := "Tahta"

	switch name {
	case "Tahta":
		fmt.Println("Hallo Tahta")
	case "Budi":
		fmt.Println("Hallo Budi")
	case "FM":
		fmt.Println("Hallo FM")

	default:
		fmt.Println("Hallo salam kenal")

	}

	//bisa menggunkan short statement

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}

}
