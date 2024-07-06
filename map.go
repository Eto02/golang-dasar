package main

import "fmt"

func main() {
	pserson := map[string]string{}
	pserson["name"] = "Tahta"
	pserson["address"] = "Lamajang"

	fmt.Println(pserson["name"])
	fmt.Println(pserson["address"])
	fmt.Println(pserson)

	//mencoba akses key yg tidak ada akan mengembalikan default value

	book := map[string]string{}
	book["title"] = "Buku Golang"
	book["author"] = "Tahta"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

}
