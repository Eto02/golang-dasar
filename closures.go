package main

import "fmt"

// Fungsi closure yang mengembalikan fungsi lain
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func resetExample() int {
	x := 0 // Variabel x diinisialisasi ulang setiap kali fungsi dipanggil
	x++
	return x
}

func main() {
	x := 10
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment()) // Output: 11
	fmt.Println(increment()) // Output: 12

	counter := makeCounter()

	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3

	add := adder()

	fmt.Println(add(10)) // Output: 10
	fmt.Println(add(20)) // Output: 30
	fmt.Println(add(30)) // Output: 60

	fmt.Println(resetExample()) // Output: 1
	fmt.Println(resetExample()) // Output: 1
	fmt.Println(resetExample()) // Output: 1
}
