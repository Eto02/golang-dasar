package main

import "fmt"

func main() {
	name := "Budi"
	if name == "Tahta" {
		fmt.Println("Hallo Tahta")
	} else if name == "Budi" {
		fmt.Println("Hallo Budi")
	} else {
		fmt.Println("Salah")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang", length)
	} else {
		fmt.Println("Nama berhasil", length)
	}
	// fmt.Println(length) //tdak bisa mengakses
}
