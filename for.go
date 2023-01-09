package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Sekarang Perulangan ke - ", counter)
	// 	counter++
	// }

	// For dengan statement

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke - ", counter)
	}

	// For Range Slice

	slice := []string{"Funny", "Violet", "Livi", "Reyhan", "Radit"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for _, value := range slice {
		fmt.Println(value)
	}

	// For Range Map

	person := make(map[string]string)
	person["name"] = "Rudy"
	person["tile"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "-", value)
	}
}
