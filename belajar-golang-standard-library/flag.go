package main

import (
	"flag"
	"fmt"
)

func main() {
	// Mendefinisikan flag
	name := flag.String("name", "World", "a name to say hello to")
	age := flag.Int("age", 30, "your age")
	verbose := flag.Bool("verbose", false, "enable verbose mode")

	// Mem-parsing flag
	flag.Parse()

	// Menggunakan nilai flag
	if *verbose {
		fmt.Println("Verbose mode enabled")
	}

	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
