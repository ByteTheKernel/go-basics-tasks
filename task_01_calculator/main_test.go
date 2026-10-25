package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "go-task-test-*")
	if err != nil {
		panic(err)
	}
	binaryName := "task"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	binary := filepath.Join(dir, binaryName)
	if err := exec.Command("go", "build", "-o", binary, "main.go").Run(); err != nil {
		os.RemoveAll(dir)
		panic(err)
	}
	testBinary = binary
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

var testBinary string

func TestProgram(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{"addition", "2 + 3\n", []string{"Результат: 5"}},
		{"subtraction", "2 - 5\n", []string{"Результат: -3"}},
		{"multiplication", "-4 * 3\n", []string{"Результат: -12"}},
		{"division", "7 / 2\n", []string{"Результат: 3.5"}},
		{"division_by_zero", "7 / 0\n", []string{"Ошибка: деление на ноль"}},
		{"unknown_operation", "7 ^ 2\n", []string{"Ошибка: неизвестная операция"}},
		{"invalid_first_number", "abc\n", []string{"Ошибка: некорректное первое число"}},
		{"invalid_second_number", "1 + abc\n", []string{"Ошибка: некорректное второе число"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(testBinary)
			cmd.Stdin = strings.NewReader(tc.input)
			var output bytes.Buffer
			cmd.Stdout = &output
			cmd.Stderr = &output
			if err := cmd.Run(); err != nil {
				t.Fatalf("program failed: %v\noutput:\n%s", err, output.String())
			}
			for _, want := range tc.expected {
				if !strings.Contains(output.String(), want) {
					t.Errorf("expected output to contain %q; got:\n%s", want, output.String())
				}
			}
		})
	}
}
