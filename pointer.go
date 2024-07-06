package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	//Pass by value
	address2 := address1
	//Pass by reference
	address3 := &address1

	address2.City = "Bandung"
	address3.City = "Lumajang"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
