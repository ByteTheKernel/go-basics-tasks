package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число")
		return
	}

	if n <= 0 {
		fmt.Println("Ошибка: N должно быть положительным")
		return
	}

	sum := 0

	for i := 0; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}

	fmt.Println("Сумма подходящих чисел:", sum)
}
