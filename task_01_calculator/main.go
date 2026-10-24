package main

import "fmt"

func main() {
	var a, b float64
	var operation string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Ошибка: некорректное первое число")
		return
	}

	fmt.Print("Введите операцию (+, -, *, /): ")
	if _, err := fmt.Scan(&operation); err != nil {
		fmt.Println("Ошибка: не удалось прочитать операцию")
		return
	}

	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Ошибка: некорректное второе число")
		return
	}

	switch operation {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		fmt.Println("Результат:", a/b)
	default:
		fmt.Println("Ошибка: неизвестная операция")
	}
}
