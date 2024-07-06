package main

import "fmt"

func main() {
	names := [...]string{"Tahta", "Failah", "Mubarak", "Budi", "Testing", "Go"}
	slice := names[4:6]
	fmt.Println(slice)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSclice := days[5:]
	fmt.Println(daysSclice)
	daysSclice[0] = "sabtu baru"
	daysSclice[1] = "minggu baru"
	fmt.Println(daysSclice)
	fmt.Println(days)

	daysSclice2 := append(daysSclice, "Libur baru")
	fmt.Println(daysSclice)
	fmt.Println(daysSclice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Tahta"
	newSlice[1] = "Tahta"
	// newSlice[2] = "Tahta" //error, harus menggunkan append

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Tahta")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "FM"

	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// jika menggunakan aappend ketika capaicty slice penuh, maka akan di buatakan array baru

	formSlice := days[:]
	toSlice := make([]string, len(formSlice), cap(formSlice))
	copy(toSlice, formSlice)

	fmt.Println(formSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
