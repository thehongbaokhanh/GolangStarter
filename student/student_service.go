package student

import (
	"fmt"
	"go-demo/ultils"
	"slices"
	"strings"
)

var StudentsList []Student

func StudentMenu() {
	for {
		ultils.ClearScreen()
		fmt.Println("========Student Management Menu========")
		fmt.Println("1. Student List")
		fmt.Println("2. Add New Student")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Search Student")
		fmt.Println("6. Exit")
		choice := ultils.ReverseInt("Enter your choice: ")

		switch choice {
		case 1:
			List()
		case 2:
			Add()
		case 3:
			Update()
		case 4:
			Delete()
		case 5:
			Search()
		case 6:
			return
		default:
			fmt.Println("Invalid choice. Please try again!")
		}
		ultils.Pause("Press Enter to continue...")
	}
}

func List() {
	fmt.Println("")
	fmt.Println("======================================List student======================================")
	if len(StudentsList) == 0 {
		fmt.Println("No student found.")
	} else {
		for _, student := range StudentsList {
			fmt.Printf("- ID: %-2d | Name: %-20s | Class: %-8s", student.Id, student.Name, student.Class)
			fmt.Print("  Score: ")
			for i, score := range student.Scores {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("| %.2f", score)
			}
			fmt.Println()
		}
	}
}

func Add() {
	Scores := []float64{}
	fmt.Println("")
	fmt.Println("========================Add New student========================")
	count := len(StudentsList)
	var id int
	if count == 0 {
		id = 1
	} else {
		id = count + 1
	}
	name := ultils.GetNonEmptyString("- Enter Name: ")
	class := ultils.GetNonEmptyString("- Enter Class: ")
	totalPoints := ultils.ReverseInt("- Enter Total Points: ")
	for i := 1; i <= totalPoints; i++ {
		for {
			point := ultils.ReverseFloat(fmt.Sprintf("- Enter Point %d: ", i))
			if point < 0 || point > 10 {
				fmt.Println("Error: Point must be between 0 and 10. Please try again.")
				continue
			}
			Scores = append(Scores, point)
			break
		}
	}

	student := Student{
		Id:     id,
		Name:   name,
		Class:  class,
		Scores: Scores,
	}
	StudentsList = append(StudentsList, student)
	fmt.Printf("Add student successfully: %v\n", student)
}

func Update() {
	fmt.Println("")
	fmt.Println("========================Update student========================")
	for {
		id := ultils.ReverseInt("Please enter the ID of the student you want to update: ")
		if id < 1 || id > len(StudentsList) {
			fmt.Println("No student found from the id you entered.")
			continue
		}
		student := GetStudentById(id)
		name := ultils.GetNonEmptyString(fmt.Sprintf("- Enter Name (%s): ", student.Name))
		class := ultils.GetNonEmptyString(fmt.Sprintf("- Enter Class (%s): ", student.Class))
		newScores := make([]float64, len(student.Scores))
		for i := 1; i <= len(student.Scores); i++ {
			point := ultils.ReverseFloat(fmt.Sprintf("- Enter Point %d (%f): ", i, student.Scores[i-1]))
			newScores[i-1] = point
		}
		student = Student{
			Id:     id,
			Name:   name,
			Class:  class,
			Scores: newScores,
		}
		StudentsList = slices.Replace(StudentsList, id-1, id, student)
		fmt.Printf("Update student successfully: %v\n", student)
		break	
	}
}

func Delete() {
	fmt.Println("")
	fmt.Println("========================Delete student========================")
	for {
		id := ultils.ReverseInt("Please enter the ID of the student you want to delete: ")
		student := GetStudentById(id)
		if len(StudentsList) < id {
			fmt.Println("No student found from the id you entered.")
			continue
		} else {
			fmt.Printf("- ID: %-2d | Name: %-20s | Class: %-8s", student.Id, student.Name, student.Class)
			fmt.Print("  Score: ")
			for i, score := range student.Scores {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("| %.2f", score)
			}
			fmt.Println()
		}
		StudentsList = slices.DeleteFunc(StudentsList, func(s Student) bool {
			return s.Id == id
		})
		fmt.Printf("Delete student successfully: %v\n", student)
		break
	}
}

func Search() {
	fmt.Println("")
	fmt.Println("========================Search for student========================")
	name := ultils.GetNonEmptyString("Please enter the name of the student you want to search: ")
	result := GetStudentByName(name)
	if len(result) == 0 {
		fmt.Println("No student found from the name you entered.")
	} else {
		for _, student := range result {
			fmt.Printf("- ID: %-2d | Name: %-20s | Class: %-8s", student.Id, student.Name, student.Class)
			fmt.Print("  Score: ")
			for i, score := range student.Scores {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("| %.2f", score)
			}
			fmt.Println()
		}
	}
}

func GetStudentById(id int) Student {
	for _, student := range StudentsList {
		if student.Id == id {
			return student
		}
	}
	return Student{}
}

func GetStudentByName(name string) []Student {
	var result []Student
	for _, student := range StudentsList {
		// case-insensitive substring match
		if strings.Contains(strings.ToLower(student.Name), strings.ToLower(name)) {
			result = append(result, student)
		}
	}
	return result
}
