package main

import "fmt"

func main() {

	counter := 0
	name := "Doni"

	// dapat merubah value variable yang ada di atas functionnya, kalau tidak ingin berubah deklarasi ulang variable dengan menggunakan :=

	increment := func() {
		name := "Rudi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
