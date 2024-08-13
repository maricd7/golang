package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("Hello, %s!\n", name)
	var year int = 2024 - age
	fmt.Printf("You were born in  %d\n", year)
}
