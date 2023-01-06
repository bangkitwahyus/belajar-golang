package main

import "fmt"

func main() {

	var name string

	name = "Miya Kazuya"
	fmt.Println(name)

	name = "Robert Blank"
	fmt.Println(name)

	var FriendName = "Gusion Alberto"
	fmt.Println(FriendName)

	var age int8 = 25
	fmt.Println(age)

	//tidak perlu menggunakan var jika sudah di dekalrasikan awal dengan :=

	country := "Indonesia"
	fmt.Println(country)

	country = "Amerika"
	fmt.Println(country)

	// Multiple Variable

	var (
		firstName = "Bobby"
		lastName  = "Saputra"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
