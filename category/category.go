package category

import (
	"encoding/json"
	"fmt"
	"os"
	"register-cli/register"
)

func ShowCategory() {
	// 카테고리 보여주기
	cli, err := loadCli("command.json")
	if err != nil {
		fmt.Println("Error loading commands:", err)
		return
	}

	for _, category := range cli.Categories {
		fmt.Println("Category:", category.Name)
		for _, cmd := range category.Commands {
			fmt.Println("  - Description:", cmd.Description)
			fmt.Println("    Command:", cmd.Command)
		}
	}
}

func loadCli(filename string) (register.Cli, error) { // 패키지명을 명시적으로 추가
	// JSON 파일에서 데이터 로드
	file, err := os.Open(filename)
	if err != nil {
		return register.Cli{}, err
	}
	defer file.Close()

	var cli register.Cli // `register.Cli`로 변경
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cli)
	if err != nil {
		return register.Cli{}, err
	}

	return cli, nil
}
