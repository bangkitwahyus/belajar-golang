package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye

	result := getGoodBye("Rian")

	fmt.Println(result)
	fmt.Println(goodbye("Rian"))
}
