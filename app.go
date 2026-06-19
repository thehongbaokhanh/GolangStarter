package main

import (
	"fmt"
	"go-demo/student"
	"go-demo/teacher"
	"go-demo/ultils"
)

func main() {
	for {
		ultils.ClearScreen()
		fmt.Println("========Management Main Menu========")
		fmt.Println("1. Student Management")
		fmt.Println("2. Teacher Management")
		fmt.Println("3. Exit")
		choice := ultils.ReverseInt("Enter your choice: ")

		switch choice {
		case 1:
			student.StudentMenu()
		case 2:
			teacher.TeacherMenu()
		case 3:
			return
		default:
			fmt.Println("Invalid choice. Please try again!")
		}
	}
}
