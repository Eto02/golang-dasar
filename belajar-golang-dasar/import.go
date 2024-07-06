package main

import (
	"belajar-golang-dasar/database"
	"belajar-golang-dasar/helper"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	result := helper.SayHello("Tahta")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye("Tahta"))

	fmt.Println(database.GetDatabase())

}
