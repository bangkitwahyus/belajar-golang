package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //Memaksa address1 datanya sama seperti address2
	address3 := &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // Memaksa semua variable yang mengarah ke struct Address berubah sesuai value yang baru

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Bali"
	fmt.Println(address4)
}
