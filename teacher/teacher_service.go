package teacher

import (
	"fmt"
	"go-demo/ultils"
	"slices"
	"strings"
)

var TeachersList []Teacher

func TeacherMenu() {
	for {
		ultils.ClearScreen()
		fmt.Println("========Teacher Management Menu========")
		fmt.Println("1. Teacher List")
		fmt.Println("2. Add New Teacher")
		fmt.Println("3. Update Teacher")
		fmt.Println("4. Delete Teacher")
		fmt.Println("5. Search Teacher")
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
	fmt.Println("=============================================List teacher=============================================")
	if len(TeachersList) == 0 {
		fmt.Println("No teacher found.")
	} else {
		for _, teacher := range TeachersList {
			fmt.Printf("- ID: %-2d | Name: %-20s | Subject: %-8s | Salary: %-8.2f | Reward: %-8.2f\n", teacher.Id, teacher.Name, teacher.Subject, teacher.Salary, teacher.Reward)
		}
	}
}

func Add() {
	fmt.Println("")
	fmt.Println("========================Add New teacher========================")
	count := len(TeachersList)
	var id int
	if count == 0 {
		id = 1
	} else {
		id = count + 1
	}
	name := ultils.GetNonEmptyString("- Enter Name: ")
	subject := ultils.GetNonEmptyString("- Enter Subject: ")
	salary := ultils.ReverseFloat("- Enter Salary: ")
	reward := ultils.ReverseFloat("- Enter Reward: ")

	teacher := Teacher{
		Id:     id,
		Name:   name,
		Subject: subject,
		Salary:  salary,
		Reward:  reward,
	}
	TeachersList = append(TeachersList, teacher)
	fmt.Printf("Add teacher successfully: %v\n", teacher)
}

func Update() {
	fmt.Println("")
	fmt.Println("========================Update teacher========================")
	for {
		id := ultils.ReverseInt("Please enter the ID of the teacher you want to update: ")
		if id < 1 || id > len(TeachersList) {
			fmt.Println("No teacher found from the id you entered.")
			continue
		}
		teacher := GetTeacherById(id)
		name := ultils.GetNonEmptyString(fmt.Sprintf("- Enter Name (%s): ", teacher.Name))
		subject := ultils.GetNonEmptyString(fmt.Sprintf("- Enter Subject (%s): ", teacher.Subject))
		salary := ultils.ReverseFloat(fmt.Sprintf("- Enter Salary (%.2f): ", teacher.Salary))
		reward := ultils.ReverseFloat(fmt.Sprintf("- Enter Reward (%.2f): ", teacher.Reward))
		teacher = Teacher{
			Id:     id,
			Name:   name,
			Subject: subject,
			Salary:  salary,
			Reward:  reward,
		}
		TeachersList = slices.Replace(TeachersList, id-1, id, teacher)
		fmt.Printf("Update teacher successfully: %v\n", teacher)
		break
	}
}

func Delete() {
	fmt.Println("")
	fmt.Println("========================Delete teacher========================")
	for {
		id := ultils.ReverseInt("Please enter the ID of the teacher you want to delete: ")
		teacher := GetTeacherById(id)
		if len(TeachersList) < id {
			fmt.Println("No teacher found from the id you entered.")
			continue
		} else {
			fmt.Printf("- ID: %-2d | Name: %-20s | Subject: %-8s | Salary: %-8.2f | Reward: %-8.2f\n", teacher.Id, teacher.Name, teacher.Subject, teacher.Salary, teacher.Reward)
			fmt.Println()
		}
		TeachersList = slices.DeleteFunc(TeachersList, func(s Teacher) bool {
			return s.Id == id
		})
		fmt.Printf("Delete teacher successfully: %v\n", teacher)
		break
	}
}

func Search() {
	fmt.Println("")
	fmt.Println("========================Search for teacher========================")
	name := ultils.GetNonEmptyString("Please enter the name of the teacher you want to search: ")
	result := GetTeacherByName(name)
	if len(result) == 0 {
		fmt.Println("No teacher found from the name you entered.")
	} else {
		for _, teacher := range result {
			fmt.Printf("- ID: %-2d | Name: %-20s | Subject: %-8s | Salary: %-8.2f | Reward: %-8.2f\n", teacher.Id, teacher.Name, teacher.Subject, teacher.Salary, teacher.Reward)
			fmt.Println()
		}
	}
}

func GetTeacherById(id int) Teacher {
	for _, teacher := range TeachersList {
		if teacher.Id == id {
			return teacher
		}
	}
	return Teacher{}
}

func GetTeacherByName(name string) []Teacher {
	var result []Teacher
	for _, teacher := range TeachersList {
		// case-insensitive substring match
		if strings.Contains(strings.ToLower(teacher.Name), strings.ToLower(name)) {
			result = append(result, teacher)
		}
	}
	return result
}
