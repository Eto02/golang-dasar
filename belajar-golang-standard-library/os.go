package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, v := range args {
		fmt.Println(v)
	}

	host, err := os.Hostname()
	if err == nil {
		fmt.Println(host)
	} else {
		fmt.Println("Error", err.Error())
	}
}
