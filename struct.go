package main

import "fmt"

// Definisi struct
type Person struct {
	Name    string
	Age     int
	Address string
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	// Membuat variabel bertipe struct Person
	var person1 Person

	// Mengisi nilai ke dalam variabel struct
	person1.Name = "John Doe"
	person1.Age = 30
	person1.Address = "123 Main St"

	// Mengakses nilai dari variabel struct
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("Address:", person1.Address)

	person2 := Person{
		Name:    "Jane Smith",
		Age:     25,
		Address: "456 Elm St",
	}
	fmt.Println(person2)
	person2.SayHello()

}
