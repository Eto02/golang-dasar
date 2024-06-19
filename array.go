package main

import (
	"fmt"
)

func main() {
	// Array nama bulan
	bulan := [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice1 := bulan[4:7]
	fmt.Println(slice1)

	slice2 := bulan[:3]
	fmt.Println(slice2)

	slice3 := bulan[3:]
	fmt.Println(slice3)

	slice4 := bulan[:]
	fmt.Println(slice4)
}
