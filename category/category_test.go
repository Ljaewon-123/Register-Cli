package category

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"register-cli/register"
)

func TestLoadCli_FileNotFound(t *testing.T) {
	_, err := loadCli("nonexistent_file.json")
	if err == nil {
		t.Error("Expected error for missing file, got nil")
	}
}

func TestLoadCli_ValidFile(t *testing.T) {
	testFile := filepath.Join("..", "testdata", "test_command.json")

	// 미리 테스트용 CLI JSON 생성
	cli := register.Cli{
		Categories: []register.Category{
			{
				Name: "Dev",
				Commands: []register.Command{
					{Description: "Echo Dev", Command: "echo dev"},
				},
			},
		},
	}
	data, _ := jsonMarshal(cli)
	_ = os.WriteFile(testFile, data, 0644)

	loaded, err := loadCli(testFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(loaded.Categories) != 1 || loaded.Categories[0].Name != "Dev" {
		t.Errorf("Invalid category loaded: %+v", loaded.Categories)
	}
}

func jsonMarshal(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
