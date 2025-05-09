package category

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"register-cli/register"

	"github.com/AlecAivazis/survey/v2"
)

func ShowCategory() {
	cli, err := loadCli("command.json")
	if err != nil {
		fmt.Println("Error loading commands:", err)
		return
	}

	if len(cli.Categories) == 0 {
		fmt.Println("카테고리가 없습니다.")
		return
	}

	// 카테고리 이름 목록 만들기
	categoryNames := []string{}
	for _, cat := range cli.Categories {
		categoryNames = append(categoryNames, cat.Name)
	}

	var selectedCategory string
	categoryPrompt := &survey.Select{
		Message: "카테고리를 선택하세요:",
		Options: categoryNames,
	}
	survey.AskOne(categoryPrompt, &selectedCategory)

	// 선택된 카테고리에서 명령어 가져오기
	var selectedCommands []register.Command
	for _, cat := range cli.Categories {
		if cat.Name == selectedCategory {
			selectedCommands = cat.Commands
			break
		}
	}

	if len(selectedCommands) == 0 {
		fmt.Println("명령어가 없습니다.")
		return
	}

	// 명령어 설명 목록 만들기
	commandOptions := []string{}
	cmdMap := make(map[string]string) // description → actual command
	for _, cmd := range selectedCommands {
		commandOptions = append(commandOptions, cmd.Description)
		cmdMap[cmd.Description] = cmd.Command
	}

	var selectedDesc string
	cmdPrompt := &survey.Select{
		Message: "실행할 명령어를 선택하세요:",
		Options: commandOptions,
	}
	survey.AskOne(cmdPrompt, &selectedDesc)

	// 실제 명령어 실행
	finalCmd := cmdMap[selectedDesc]
	fmt.Printf("\n[실행 중] %s\n\n", finalCmd)
	runCommand(finalCmd)
}

func loadCli(filename string) (register.Cli, error) {
	file, err := os.Open(filename)
	if err != nil {
		return register.Cli{}, err
	}
	defer file.Close()

	var cli register.Cli
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cli)
	if err != nil {
		return register.Cli{}, err
	}

	return cli, nil
}

func runCommand(command string) {
	cmd := exec.Command("sh", "-c", command) // macOS/Linux
	// cmd := exec.Command("cmd", "/C", command) // Windows용일 경우

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("명령어 실행 실패: %v\n", err)
	}
}
