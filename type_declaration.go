package main

import "fmt"

func main() {

	type noKtp string
	ktpKu := "11111111111"
	contoh := "2222222222"
	contohKtp := noKtp(contoh)
	fmt.Println(ktpKu)
	fmt.Println(contohKtp)

}
