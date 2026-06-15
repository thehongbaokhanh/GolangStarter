package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	var name string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&name)
// 	fmt.Printf("Hello, %s!\n", name)
// }

func main() {
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name := scanner.Text()
		fmt.Printf("Welcome to the Go programming language, %s!\n", name)
	}
}
