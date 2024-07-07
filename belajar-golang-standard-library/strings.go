package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contoh penggunaan strings.Contains
	str := "Hello, world!"
	substr := "world"
	fmt.Println("Contains:", strings.Contains(str, substr)) // Output: true

	// Contoh penggunaan strings.Count
	str2 := "banana"
	substr2 := "a"
	fmt.Println("Count:", strings.Count(str2, substr2)) // Output: 3

	// Contoh penggunaan strings.HasPrefix dan strings.HasSuffix
	fmt.Println("HasPrefix:", strings.HasPrefix(str, "Hello"))  // Output: true
	fmt.Println("HasSuffix:", strings.HasSuffix(str, "world!")) // Output: true

	// Contoh penggunaan strings.Index dan strings.LastIndex
	fmt.Println("Index:", strings.Index(str, "o"))         // Output: 4
	fmt.Println("LastIndex:", strings.LastIndex(str, "o")) // Output: 8

	// Contoh penggunaan strings.Join
	slice := []string{"foo", "bar", "baz"}
	separator := ", "
	fmt.Println("Join:", strings.Join(slice, separator)) // Output: foo, bar, baz

	// Contoh penggunaan strings.Split
	str3 := "a,b,c"
	separator2 := ","
	fmt.Println("Split:", strings.Split(str3, separator2)) // Output: [a b c]

	// Contoh penggunaan strings.Replace
	old := "world"
	new := "Go"
	fmt.Println("Replace:", strings.Replace(str, old, new, -1)) // Output: Hello, Go!

	// Contoh penggunaan strings.ToLower dan strings.ToUpper
	fmt.Println("ToLower:", strings.ToLower(str)) // Output: hello, world!
	fmt.Println("ToUpper:", strings.ToUpper(str)) // Output: HELLO, WORLD!

	// Contoh penggunaan strings.Trim, strings.TrimSpace, strings.TrimPrefix, dan strings.TrimSuffix
	str4 := "  Hello, world!  "
	fmt.Println("TrimSpace:", strings.TrimSpace(str4))         // Output: Hello, world!
	fmt.Println("Trim:", strings.Trim(str4, " !"))             // Output: Hello, world
	fmt.Println("TrimPrefix:", strings.TrimPrefix(str4, "  ")) // Output: Hello, world!
	fmt.Println("TrimSuffix:", strings.TrimSuffix(str4, "  ")) // Output:   Hello, world!

	// Contoh penggunaan strings.Repeat
	str5 := "Go"
	n := 3
	fmt.Println("Repeat:", strings.Repeat(str5, n)) // Output: GoGoGo
}
