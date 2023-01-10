package main

import "fmt"

// Memanggil function ke main function

func sayHello() {

	names := "Rani"

	fmt.Println("Hello", names)
}

func main() {

	for i := 0; i < 5; i++ {
		sayHello()
	}
}
