package register

import (
	"encoding/json"
	"fmt"
	"os"
)

// 명령어 구조체
type Command struct {
	Description string `json:"description"`
	Command     string `json:"command"`
}

// 카테고리 구조체
type Category struct {
	Name     string    `json:"name"`
	Commands []Command `json:"commands"`
}

type Cli struct {
	Categories []Category `json:"categories"`
}

// JSON 파일에서 데이터를 불러오는 함수
func loadCli(filename string) (Cli, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Cli{}, nil // 파일이 없으면 빈 구조체 반환
	}
	defer file.Close()

	var cli Cli
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cli)
	if err != nil {
		return Cli{}, err
	}

	return cli, nil
}

// JSON 파일에 데이터를 저장하는 함수
func saveCli(filename string, cli Cli) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(cli)
}

// 새로운 Command를 추가 또는 수정하는 함수
func saveCommand(filename string, categoryName string, command Command) error {
	// 기존 데이터 불러오기
	cli, err := loadCli(filename)
	if err != nil {
		return err
	}

	// 해당 카테고리가 존재하는지 확인
	var categoryIndex int = -1
	for i, cat := range cli.Categories {
		if cat.Name == categoryName {
			categoryIndex = i
			break
		}
	}

	if categoryIndex == -1 {
		// 카테고리가 없으면 새로 추가
		cli.Categories = append(cli.Categories, Category{
			Name:     categoryName,
			Commands: []Command{command},
		})
	} else {
		// 기존 카테고리 내부에서 명령어가 존재하는지 확인 후 수정 또는 추가
		found := false
		for j, cmd := range cli.Categories[categoryIndex].Commands {
			if cmd.Description == command.Description {
				cli.Categories[categoryIndex].Commands[j] = command
				found = true
				break
			}
		}

		if !found {
			// 기존 명령어가 없으면 새로 추가
			cli.Categories[categoryIndex].Commands = append(cli.Categories[categoryIndex].Commands, command)
		}
	}

	// 업데이트된 데이터를 저장
	return saveCli(filename, cli)
}

// RegisterCommand 실행 함수
func RegisterCommand(name string, description string, command string) {
	err := saveCommand("command.json", name, Command{
		Description: description,
		Command:     command,
	})
	if err != nil {
		fmt.Println("Error saving command:", err)
		return
	}

	// 저장된 데이터 확인
	cli, err := loadCli("command.json")
	if err != nil {
		fmt.Println("Error loading commands:", err)
		return
	}

	// 결과 출력
	for _, category := range cli.Categories {
		fmt.Println("Category:", category.Name)
		for _, cmd := range category.Commands {
			fmt.Println("  - Description:", cmd.Description)
			fmt.Println("    Command:", cmd.Command)
		}
	}
}
