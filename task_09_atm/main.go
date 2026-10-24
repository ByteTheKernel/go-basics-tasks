package main

import "fmt"

func main() {
	var balance float64

	fmt.Print("Введите начальный баланс: ")
	if _, err := fmt.Scan(&balance); err != nil {
		fmt.Println("Ошибка: необходимо ввести число")
		return
	}

	if balance < 0 {
		fmt.Println("Ошибка: баланс не может быть отрицательным")
		return
	}

	var choice int

	fmt.Println("\nВыберите действие:")
	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Пополнить баланс")
	fmt.Println("3. Снять средства")

	fmt.Print("Ваш выбор: ")
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("Ошибка: некорректный выбор")
		return
	}

	switch choice {
	case 1:
		fmt.Printf("Текущий баланс: %.2f руб.\n", balance)

	case 2:
		var amount float64

		fmt.Print("Введите сумму пополнения: ")
		if _, err := fmt.Scan(&amount); err != nil {
			fmt.Println("Ошибка: некорректная сумма")
			return
		}

		if amount <= 0 {
			fmt.Println("Ошибка: сумма должна быть положительной")
			return
		}

		balance += amount

		fmt.Printf("Баланс успешно пополнен.\n")
		fmt.Printf("Текущий баланс: %.2f руб.\n", balance)

	case 3:
		var amount float64

		fmt.Print("Введите сумму снятия: ")
		if _, err := fmt.Scan(&amount); err != nil {
			fmt.Println("Ошибка: некорректная сумма")
			return
		}

		if amount <= 0 {
			fmt.Println("Ошибка: сумма должна быть положительной")
			return
		}

		if amount > balance {
			fmt.Println("Ошибка: недостаточно средств")
			return
		}

		balance -= amount

		fmt.Println("Средства успешно сняты.")
		fmt.Printf("Текущий баланс: %.2f руб.\n", balance)

	default:
		fmt.Println("Ошибка: неизвестная операция")
	}
}
