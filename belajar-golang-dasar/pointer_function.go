package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	alamat := Address{}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}
