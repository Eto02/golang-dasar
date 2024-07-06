package main

import "fmt"

func main() {

	sayHello()
	greet("Tahta")
	fmt.Println(add(2, 5))
	fmt.Println(divide(2, 5))
	fmt.Println(split(10))
	fmt.Println(sum(1, 23, 4, 4, 5))
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))
	fmt.Println(Rectangle{height: 5, width: 2}.Area())

	newAdd := func(a int, b int) int {
		return a + b
	}

	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Println("3 + 4 =", compute(newAdd, 3, 4))   // Output: 3 + 4 = 7
	fmt.Println("3 * 4 =", compute(multiply, 3, 4)) // Output: 3 * 4 = 12

	fmt.Println("5! =", factorial(5))
}

func sayHello() {
	fmt.Println("Hello")
}

func greet(name string) {
	fmt.Printf("Hello, %s\n", name)
}
func add(a int, b int) int {
	return a + b
}
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func compute(operation func(int, int) int, a int, b int) int {
	return operation(a, b)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
