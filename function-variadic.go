package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {

	total := sumAll(10, 50, 20, 90)
	fmt.Println("Hasil Penjumlahan ", total)

	// Slice Parameter

	slice := []int{10, 40, 60, 70}
	total = sumAll(slice...)
	fmt.Println("Hasil Penjumlahan ", total)

}
