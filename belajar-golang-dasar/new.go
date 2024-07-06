package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	alamat1 := new(Address)
	alamat2 := new(Address)
	alamat2.City = "Lumajang"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
