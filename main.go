package main

import (
	"bufio"
	"fmt"
	"os"
	"register-cli/category"
	"register-cli/register"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose an option:")
	fmt.Println("1. Show category")
	fmt.Println("2. Register command")
	fmt.Println("3. Exit")

	var choice int
	fmt.Print("Select number: ")
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("This wrong number")
		return
	}

	// 개행 문자 제거 (버퍼 정리)
	reader.ReadString('\n')

	switch choice {
	case 1:
		fmt.Println("Show category")
		category.ShowCategory()
	case 2:
		createNewCommand(reader)
	case 3:
		fmt.Println("exit")
	default:
		fmt.Println("Wrong number.")
	}
}

func createNewCommand(reader *bufio.Reader) {
	fmt.Println("Register command")

	fmt.Print("Enter the category to register: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category) // 개행 문자 제거

	fmt.Print("Enter command description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Enter command: ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

	register.RegisterCommand(category, description, command)
}
