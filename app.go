package main

import "fmt"

// Bài tập 1: Viết chương trình in ra các số từ 1 đến 100, nhưng bỏ qua các số 6, 48, 75 và 89. Các số được in ra sẽ được phân cách bằng dấu phẩy và một khoảng trắng. Sau số cuối cùng (100) sẽ không có dấu phẩy.

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i == 6|| i == 48 || i == 75 || i == 89 {
// 			continue
// 		}else if i == 100 {
// 			fmt.Print(i , ".")
// 			break
// 		}else {
// 		fmt.Print(i , ", ")
// 		}
// 	}
// }

// Bài tập 2: Viết chương trình in ra tất cả các số lẻ từ 1 đến 100. Các số được in ra sẽ được phân cách bằng dấu phẩy và một khoảng trắng. Sau số cuối cùng (99) sẽ không có dấu phẩy.

// func main() {
// 	var count int = 1
// 	for i := 1; i <= 100; i++ {
// 		if i % 2 == 0 {
// 			continue
// 		}else if i == 100 {
// 			fmt.Print(i , ".")
// 			break
// 		}else {
// 			if count < 3{
// 				fmt.Print(i , ", ")
// 				count++
// 			}else {
// 				fmt.Print(i , "\n")
// 				count = 1
// 			}
// 		}
// 	}
// }

func main() {
	var n1 int = -1
	var n2 int = -1

	fmt.Print("Enter the number of the first multiplication table: ")
	fmt.Scanln(&n1)

	fmt.Print("Enter the number of the second multiplication table: ")
	fmt.Scanln(&n2)

	if n1 < 1 || n1 > 10 {
		fmt.Println("Please enter a number between 1 and 10.")
		fmt.Print("Enter the number of the first multiplication table: ")
		fmt.Scanln(&n1)
	}
	if n2 < 1 || n2 > 10 {
		fmt.Println("Please enter a number between 1 and 10.")
		fmt.Print("Enter the number of the second multiplication table: ")
		fmt.Scanln(&n2)
	}

	fmt.Printf("Multiplication table of %d:\n", n1)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n1, i, n1*i)
	}

	fmt.Printf("Multiplication table of %d:\n", n2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n2, i, n2*i)
	}

}
