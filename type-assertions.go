package main

import "fmt"

func random() interface{} {
	return "Hai"
}

func main() {
	result := random()

	// Type Assertion

	// resultString := result.(string)
	// fmt.Println(resultString)

	// Handle with switch

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Integer")
	default:
		fmt.Println("Unknown type")
	}
}
