package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	roni := Man{"Roni"}
	roni.Married()

	fmt.Println(roni)
}
