package main

import "fmt"

// Type declaration, untuk membuat alias agar lebih singkat

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "******"
	} else {
		return name
	}
}

func main() {

	sayHelloWithFilter("Funi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}
