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
	fmt.Print("옵션 번호 입력: ")
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("This wrong number")
		return
	}

	// 개행 문자 제거 (버퍼 정리)
	reader.ReadString('\n')

	switch choice {
	case 1:
		fmt.Println("카테고리 보여주기")
		category.ShowCategory()
	case 2:
		createNewCommand(reader)
	case 3:
		fmt.Println("종료하기")
	default:
		fmt.Println("잘못된 번호입니다.")
	}
}

func createNewCommand(reader *bufio.Reader) {
	fmt.Println("명령어 등록하기")

	fmt.Print("등록할 카테고리 입력: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category) // 개행 문자 제거

	fmt.Print("명령어 설명 입력: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("명령어 입력: ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

	register.RegisterCommand(category, description, command)
}
