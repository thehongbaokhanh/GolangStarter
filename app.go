package main

import "fmt"

// Bài 1:
// func sum(a int) int {
// 	result := 0
// 	if a > 0 {
// 		result = a + sum(a-1)
// 	}

// 	return result
// }



// Bài 2:

// func fibonacci(count int)  int {
// 	a := 0
// 	b := 1
// 	result := 0
// 	sumResult := 0
// 	for i := 0; i < count; i++ {
// 		result = a + b
// 		a = b
// 		b = result
// 		sumResult = sum(result, sumResult)
// 		fmt.Println( result)
// 	}
// 	return sumResult
// }

// func sum(a, b int) int {
// 	result := 0
// 	result = a + b
// 	return result
// }


func main() {
	var count int
	fmt.Print("Nhập số lượng phần tử Fibonacci cần tính: ")
	fmt.Scanln(&count)
	sumFib := fibonacci(count)
	fmt.Printf("Tổng của các số Fibonacci từ 0 đến %d là: %d\n", count, sumFib)
}
