package main

import "fmt"

func random() any {
	return "OK"
}
func typeSwitch(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("Ini adalah sebuah integer:", v)
	case string:
		fmt.Println("Ini adalah sebuah string:", v)
	default:
		fmt.Println("Tipe variabel tidak dikenali")
	}
}
func main() {
	var i any = random()

	// Type assertion
	s, ok := i.(string)
	if ok {
		fmt.Println("Value is string:", s)
	} else {
		fmt.Println("Value is not a string")
	}

	f, ok := i.(float64)
	if ok {
		fmt.Println("Value is float64:", f)
	} else {
		fmt.Println("Value is not a float64")
	}
	typeSwitch(i)

}
