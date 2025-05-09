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
		fmt.Println("Not found category.")
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
		fmt.Println("Category selection error:", err)
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
		fmt.Println("Not found command.")
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
		fmt.Println("Command selection error:", err)
		return
	}

	if len(selectedDisplays) == 0 {
		fmt.Println("Not found command.")
		return
	}

	for _, display := range selectedDisplays {
		cmdStr := displayToCmd[display]
		fmt.Printf("\n[==running==] %s\n", cmdStr)
		runCommand(cmdStr)
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

func runCommand(command string) {
	cmd := exec.Command("sh", "-c", command) // macOS/Linux
	// cmd := exec.Command("cmd", "/C", command) // Windows용일 경우

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Command running is failed: %v\n", err)
	}
}
