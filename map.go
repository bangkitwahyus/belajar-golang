package main

import "fmt"

func main() {

	// MAP

	person := map[string]string{
		"name":    "Denny",
		"address": "Surabaya",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	// Function make and delete

	book := make(map[string]string)
	book["title"] = "Practice Go Lang"
	book["author"] = "Roberto Santos"
	book["wrong"] = "Ups"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))

}
