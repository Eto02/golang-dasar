package main

import "fmt"

func main() {

	a := 80
	b := 60
	lulus := a > 75
	gagal := b > 75

	benar := lulus && gagal

	fmt.Println(lulus)
	fmt.Println(gagal)
	fmt.Println(benar)

}
