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
		{"positive_even_divisible_by_3_and_5", "30\n", []string{"Число положительное", "Число чётное", "Число делится на 3", "Число делится на 5", "Число делится одновременно на 3 и на 5"}},
		{"negative_odd", "-7\n", []string{"Число отрицательное", "Число нечётное", "Число не делится на 3", "Число не делится на 5", "Число не делится одновременно на 3 и на 5"}},
		{"zero", "0\n", []string{"Число равно нулю", "Число чётное", "Число делится одновременно на 3 и на 5"}},
		{"divisible_only_by_3", "9\n", []string{"Число делится на 3", "Число не делится на 5"}},
		{"divisible_only_by_5", "10\n", []string{"Число не делится на 3", "Число делится на 5"}},
		{"invalid_input", "x\n", []string{"Ошибка: необходимо ввести целое число"}},
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
