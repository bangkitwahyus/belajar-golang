package main

import "fmt"

func main() {

	// Switch

	var name = "Lidya"

	switch name {
	case "Lia":
		fmt.Println("Hallo Lia")
	case "Lina":
		fmt.Println("Hallo Lina")
	case "Lidya":
		fmt.Println("Hallo Lidya")
	default:
		fmt.Println("Hallo Kenalan Donk")
	}

	// Short statement Switch

	// switch lenght := len(name); lenght > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	// Switch tanda kondisi

	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Nama Terlalu Panjang")
	case lenght > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
