package main

import "fmt"

func main() {

	// Break , untuk menghentikan paksa sebuah perulangan sesuai kondisi yang ditentukan

	for i := 0; i < 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke -", i)

	}
}
