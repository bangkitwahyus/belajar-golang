package main

import "fmt"

func main() {

	// Slice Array

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Append

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Libur")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Make Slice

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Funny"
	newSlice[1] = "Violet"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy Slice

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// Slice dan Array

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
