package main

import "fmt"

func main() {

	// Continue , melanjutkan perulangan sesuai kondisi yang diiinginkan

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke-", i)
	}
}
