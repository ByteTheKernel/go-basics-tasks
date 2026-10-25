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
		{"view_balance", "100\n1\n", []string{"Текущий баланс: 100.00 руб."}},
		{"deposit", "100\n2\n50\n", []string{"Баланс успешно пополнен.", "Текущий баланс: 150.00 руб."}},
		{"withdraw", "100\n3\n40\n", []string{"Средства успешно сняты.", "Текущий баланс: 60.00 руб."}},
		{"withdraw_all", "100\n3\n100\n", []string{"Текущий баланс: 0.00 руб."}},
		{"insufficient_funds", "100\n3\n101\n", []string{"Ошибка: недостаточно средств"}},
		{"zero_deposit", "100\n2\n0\n", []string{"Ошибка: сумма должна быть положительной"}},
		{"negative_withdraw", "100\n3\n-1\n", []string{"Ошибка: сумма должна быть положительной"}},
		{"negative_balance", "-10\n", []string{"Ошибка: баланс не может быть отрицательным"}},
		{"unknown_operation", "100\n9\n", []string{"Ошибка: неизвестная операция"}},
		{"invalid_amount", "100\n2\nabc\n", []string{"Ошибка: некорректная сумма"}},
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
