package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter the name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello", name)
}