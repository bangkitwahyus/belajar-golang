package main

import "fmt"

func main() {

	//Type Declaration

	type NoKTP string
	type Married bool

	var noKTPDenny NoKTP = "12362716486127"
	var marriedStatus Married = true

	fmt.Println(noKTPDenny)
	fmt.Println(marriedStatus)
}
