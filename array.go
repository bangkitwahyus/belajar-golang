package main

import "fmt"

func main() {

	// Array Manual

	var names [3]string

	names[0] = "Funny"
	names[1] = "violet"
	names[2] = "Fajria"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Array Langsung

	var values = [3]int8{
		80, 10, 90,
	}

	fmt.Println(values)

	// Untuk Mengetahui panjang array bisa lenggunakan len(array)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}
