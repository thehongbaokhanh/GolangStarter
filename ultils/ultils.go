package ultils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Pause(prompt string) string {
	fmt.Println("")
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ReverseInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		reversed, err := strconv.Atoi(input)
		if err == nil && reversed > 0 {
			return reversed
		}
		fmt.Println("Error: Invalid input. Please enter a valid integer.")

	}
}


func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}
		fmt.Println("Error: Non-empty input required. Please try again.")

	}
}

func ReverseFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		reversed, err := strconv.ParseFloat	(input, 64)
		if err == nil && reversed > 0 {
			return reversed
		}
		fmt.Println("Error: Invalid input. Please enter a valid float.")

	}
}

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
