package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Bài 1:
func sumOfNumber(a int) int {
	result := 0
	if a > 0 {
		result = a + sumOfNumber(a-1)
	}

	return result
}

// menu :

func fibonacci(count int) int {
	a := 0
	b := 1
	result := 0
	sumResult := 0
	for i := 0; i < count; i++ {
		result = a + b
		a = b
		b = result
		sumResult = sum(result, sumResult)
		// fmt.Println(result)
	}
	return sumResult
}

// readInt reads an integer from standard input after displaying the prompt
func readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return 0, fmt.Errorf("empty input")
	}
	n, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func sum(a, b int) int {
	result := 0
	result = a + b
	return result
}

func main() {
	for {
		fmt.Println("=-=-=-=-=-=-=-=Menu-=-=-=-=-=-=-=-=")
		fmt.Println("[1] Tính tổng các số từ 1 đến n")
		fmt.Println("[2] Tính tổng dãy Fibonacci")
		fmt.Println("[0] Thoát")
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		choice, err := readInt("Nhập lựa chọn của bạn: ")
		if err != nil || choice < 0 || choice > 2 {
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
			continue
		}

		switch choice {
		case 1:
			n, err := readInt("Nhập số n: ")
			if err != nil || n < 0 {
				fmt.Println("Số n phải là một số nguyên dương. Vui lòng thử lại.")
				continue
			}
			result := sumOfNumber(n)
			fmt.Printf("Tổng các số từ 1 đến %d là: %d\n", n, result)
		case 2:
			count, err := readInt("Nhập số lượng phần tử Fibonacci: ")
			if err != nil || count < 0 {
				fmt.Println("Số lượng phần tử phải là một số nguyên dương. Vui lòng thử lại.")
				continue
			}
			result := fibonacci(count)
			fmt.Printf("Tổng dãy Fibonacci với %d phần tử là: %d\n", count, result)
		case 0:
			fmt.Println("Thoát chương trình. Tạm biệt!")
			return
		}
	}
}
