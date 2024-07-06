package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func printShapeInfo(shape Shape) {
	fmt.Println("Area:", shape.Area())
	fmt.Println("Perimeter:", shape.Perimeter())
}

func main() {
	person := Person{Name: "Tahta"}
	sayHello(person)

	circle := Circle{Radius: 5}
	printShapeInfo(circle)

}
