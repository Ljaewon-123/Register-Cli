package register

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRegisterCommand(t *testing.T) {
	testFile := filepath.Join("..", "testdata", "test_command.json")
	os.Remove(testFile) // 기존 테스트 파일 제거 (클린 테스트용)

	// 테스트용 커맨드 등록
	err := saveCommand(testFile, "TestCategory", Command{
		Description: "test description",
		Command:     "echo test",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// 저장된 결과 확인
	cli, err := loadCli(testFile)
	if err != nil {
		t.Fatalf("Failed to load CLI data: %v", err)
	}

	// 카테고리 및 명령어 유효성 검사
	if len(cli.Categories) != 1 {
		t.Fatalf("Expected 1 category, got %d", len(cli.Categories))
	}

	cat := cli.Categories[0]
	if cat.Name != "TestCategory" {
		t.Errorf("Expected category name 'TestCategory', got '%s'", cat.Name)
	}

	if len(cat.Commands) != 1 {
		t.Fatalf("Expected 1 command, got %d", len(cat.Commands))
	}

	cmd := cat.Commands[0]
	if cmd.Description != "test description" || cmd.Command != "echo test" {
		t.Errorf("Command mismatch: got %+v", cmd)
	}
}

func TestLoadCli_EmptyFile(t *testing.T) {
	testFile := filepath.Join("..", "testdata", "empty.json")
	_ = os.WriteFile(testFile, []byte(""), 0644)

	_, err := loadCli(testFile)
	if err == nil {
		t.Error("Expected error for empty file, got nil")
	}
}

func TestSaveCli(t *testing.T) {
	testFile := filepath.Join("..", "testdata", "save_test.json")

	cli := Cli{
		Categories: []Category{
			{
				Name: "Sample",
				Commands: []Command{
					{Description: "Say Hello", Command: "echo hello"},
				},
			},
		},
	}

	err := saveCli(testFile, cli)
	if err != nil {
		t.Fatalf("Failed to save CLI: %v", err)
	}

	// 파일 확인
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}

	var loaded Cli
	if err := json.Unmarshal(data, &loaded); err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}
}
