package category

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"register-cli/register"

	"github.com/AlecAivazis/survey/v2"
)

// 에러 정의
var (
	errNotFoundCategory  = errors.New("category not found")
	errNotFoundCommand   = errors.New("command not found")
	errCategorySelection = errors.New("category selection error")
	errCommandSelection  = errors.New("command selection error")
	errCommandExecution  = errors.New("command execution failed")
)

func ShowCategory() {
	cli, err := loadCli("command.json")
	if err != nil {
		// 에러 반환
		handleError(err)
		return
	}

	if len(cli.Categories) == 0 {
		// 카테고리 없음
		handleError(errNotFoundCategory)
		return
	}

	// 카테고리 이름 목록 만들기
	categoryNames := make([]string, len(cli.Categories))
	for i, cat := range cli.Categories {
		categoryNames[i] = cat.Name
	}

	var selectedCategory string
	categoryPrompt := &survey.Select{
		Message: "Please select a category:",
		Options: categoryNames,
	}
	if err := survey.AskOne(categoryPrompt, &selectedCategory); err != nil {
		// 카테고리 선택 오류
		handleError(errCategorySelection)
		return
	}

	// 선택된 카테고리에서 명령어 가져오기
	var selectedCommands []register.Command
	for _, cat := range cli.Categories {
		if cat.Name == selectedCategory {
			selectedCommands = cat.Commands
			break
		}
	}

	if len(selectedCommands) == 0 {
		// 명령어 없음
		handleError(errNotFoundCommand)
		return
	}

	// 표시용 옵션 만들기 + 매핑
	commandOptions := []string{}
	displayToCmd := make(map[string]string)

	for _, cmd := range selectedCommands {
		display := fmt.Sprintf("%s\n  :: description: %s", cmd.Command, cmd.Description)
		commandOptions = append(commandOptions, display)
		displayToCmd[display] = cmd.Command
	}

	var selectedDisplays []string
	cmdPrompt := &survey.MultiSelect{
		Message: "Select multiple commands to run (select with spacebar):",
		Options: commandOptions,
	}
	if err := survey.AskOne(cmdPrompt, &selectedDisplays); err != nil {
		// 명령어 선택 오류
		handleError(errCommandSelection)
		return
	}

	if len(selectedDisplays) == 0 {
		// 명령어 없음
		handleError(errNotFoundCommand)
		return
	}

	for _, display := range selectedDisplays {
		cmdStr := displayToCmd[display]
		fmt.Printf("\n[==running==] %s\n", cmdStr)
		err := runCommand(cmdStr)
		if err != nil {
			handleError(err)
			return
		}
	}
}

// 에러 처리 함수
func handleError(err error) {
	if err != nil {
		fmt.Printf("[Error] %v\n", err) // 오류 메시지 출력
	}
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

func runCommand(command string) error {
	cmd := exec.Command("sh", "-c", command) // macOS/Linux
	// cmd := exec.Command("cmd", "/C", command) // Windows용일 경우

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %v", errCommandExecution, err) // 명령어 실행 실패 오류 반환
	}
	return nil
}
