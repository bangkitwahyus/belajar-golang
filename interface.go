package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHai(hasName HasName) {
	fmt.Println("Hello ", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {

	dodi := Person{Name: "Dodi"}
	sayHai(dodi)

	cat := Animal{Name: "Bocil"}
	sayHai(cat)

}
