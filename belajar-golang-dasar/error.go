package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembangi int) (int, error) {
	if pembangi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembangi, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err)
	}
}
