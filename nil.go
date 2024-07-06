package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	data := NewMap("Tahta")

	if data == nil {
		fmt.Println("Data map kosong")
	} else {
		fmt.Println(data["name"])
	}
}
