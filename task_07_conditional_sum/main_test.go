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
		{"below_six", "5\n", []string{"Сумма подходящих чисел: 0"}},
		{"six", "6\n", []string{"Сумма подходящих чисел: 6"}},
		{"twenty", "20\n", []string{"Сумма подходящих чисел: 36"}},
		{"thirty", "30\n", []string{"Сумма подходящих чисел: 90"}},
		{"zero", "0\n", []string{"Ошибка: N должно быть положительным"}},
		{"invalid", "x\n", []string{"Ошибка: необходимо ввести целое число"}},
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
