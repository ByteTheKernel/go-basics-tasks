package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите целое число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число")
		return
	}

	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно нулю")
	}

	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	if n%3 == 0 {
		fmt.Println("Число делится на 3")
	} else {
		fmt.Println("Число не делится на 3")
	}

	if n%5 == 0 {
		fmt.Println("Число делится на 5")
	} else {
		fmt.Println("Число не делится на 5")
	}

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Число делится одновременно на 3 и на 5")
	} else {
		fmt.Println("Число не делится одновременно на 3 и на 5")
	}
}
