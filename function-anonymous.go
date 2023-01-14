package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked ! ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

// func blockAdmin(name string) bool{
// 	return name == "Admin"
// }
// func blockUser(name string) bool {
// 	return name == "User"
// }

func main() {

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Randy", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("Randy", func(name string) bool {
		return name == "root"
	})
}