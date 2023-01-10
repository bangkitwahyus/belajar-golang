package main

import "fmt"

func getFullName() (string, string) {
	return "Funi", "Cintya"
}

func main() {

	//  tanda _ digunakan untuk ignore return value yang tidak dibutuhkan

	firstName, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(lastName)
}
