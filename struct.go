package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var doni Customer
	doni.Name = "Doni"
	doni.Address = "Surabaya"
	doni.Age = 27

	fmt.Println(doni.Name)
	fmt.Println(doni.Address)
	fmt.Println(doni.Age)

	// Struct Literals

	joko := Customer{
		Name:    "Joko",
		Address: "Jogjakarta",
		Age:     25,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 29}
	fmt.Println(budi)
}
