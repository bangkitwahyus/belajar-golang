package main

import "fmt"

func main() {

	// IF

	name := "Doniawan"

	if name == "Doni" {
		fmt.Println("Hallo Doni")
	} else if name == "Dina" {
		fmt.Println("Hallo Dina")
	} else if name == "Dini" {
		fmt.Println("Hallo Dini")
	} else {
		fmt.Println("Hi Kenalan Donk")
	}

	// Short Statement di IF

	if lenght := len(name); lenght > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
