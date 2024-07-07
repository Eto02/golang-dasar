package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Contoh penggunaan strconv.Atoi
	intStr := "123"
	intValue, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Atoi:", intValue) // Output: 123
	}

	// Contoh penggunaan strconv.Itoa
	intValue = 456
	intStr = strconv.Itoa(intValue)
	fmt.Println("Itoa:", intStr) // Output: "456"

	// Contoh penggunaan strconv.ParseFloat
	floatStr := "123.45"
	floatValue, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("ParseFloat:", floatValue) // Output: 123.45
	}

	// Contoh penggunaan strconv.FormatFloat
	floatValue = 678.90
	floatStr = strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Println("FormatFloat:", floatStr) // Output: "678.90"

	// Contoh penggunaan strconv.ParseBool
	boolStr := "true"
	boolValue, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
	} else {
		fmt.Println("ParseBool:", boolValue) // Output: true
	}

	// Contoh penggunaan strconv.FormatBool
	boolValue = false
	boolStr = strconv.FormatBool(boolValue)
	fmt.Println("FormatBool:", boolStr) // Output: "false"

	// Contoh penggunaan strconv.ParseInt
	intStr = "789"
	intValue64, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
	} else {
		fmt.Println("ParseInt:", intValue64) // Output: 789
	}

	// Contoh penggunaan strconv.FormatInt
	intValue64 = 123
	intStr = strconv.FormatInt(intValue64, 10)
	fmt.Println("FormatInt:", intStr) // Output: "123"

	// Contoh penggunaan strconv.ParseUint
	uintStr := "456"
	uintValue64, err := strconv.ParseUint(uintStr, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to uint64:", err)
	} else {
		fmt.Println("ParseUint:", uintValue64) // Output: 456
	}

	// Contoh penggunaan strconv.FormatUint
	uintValue64 = 789
	uintStr = strconv.FormatUint(uintValue64, 10)
	fmt.Println("FormatUint:", uintStr) // Output: "789"
}
