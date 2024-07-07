package main

import (
	"fmt"
	"math"
)

func main() {
	// Contoh penggunaan math.Abs (menghitung nilai absolut)
	x := -42.7
	fmt.Println("Abs:", math.Abs(x)) // Output: 42.7

	// Contoh penggunaan math.Sqrt (menghitung akar kuadrat)
	y := 16.0
	fmt.Println("Sqrt:", math.Sqrt(y)) // Output: 4.0

	// Contoh penggunaan math.Pow (menghitung pangkat)
	base := 2.0
	exponent := 3.0
	fmt.Println("Pow:", math.Pow(base, exponent)) // Output: 8.0

	// Contoh penggunaan math.Max dan math.Min (mencari nilai maksimum dan minimum)
	a := 7.5
	b := 3.8
	fmt.Println("Max:", math.Max(a, b)) // Output: 7.5
	fmt.Println("Min:", math.Min(a, b)) // Output: 3.8

	// Contoh penggunaan fungsi trigonometri (Sin, Cos, Tan)
	angle := math.Pi / 4                 // 45 derajat
	fmt.Println("Sin:", math.Sin(angle)) // Output: 0.7071067811865476
	fmt.Println("Cos:", math.Cos(angle)) // Output: 0.7071067811865476
	fmt.Println("Tan:", math.Tan(angle)) // Output: 1.0

	// Contoh penggunaan math.Round, math.Floor, dan math.Ceil (pembulatan angka)
	value := 3.7
	fmt.Println("Round:", math.Round(value)) // Output: 4
	fmt.Println("Floor:", math.Floor(value)) // Output: 3
	fmt.Println("Ceil:", math.Ceil(value))   // Output: 4

}
