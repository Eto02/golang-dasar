package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}
func main() {
	tahta := Man{"Tahta"}
	tahta.Married()
	fmt.Println(tahta.Name)
}
