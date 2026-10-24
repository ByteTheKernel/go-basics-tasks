package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число")
		return
	}

	if n < 0 {
		fmt.Println("Ошибка: N должно быть положительным")
		return
	}

	positiveCount := 0
	negativeCount := 0
	zeroCount := 0

	for i := 0; i < n; i++ {
		var number int

		fmt.Printf("Введите число %d: ", i+1)
		if _, err := fmt.Scan(&number); err != nil {
			fmt.Println("Ошибка: необходимо ввести целое число")
			return
		}

		if number > 0 {
			positiveCount++
		} else if number < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	fmt.Println("Количество положительных чисел:", positiveCount)
	fmt.Println("Количество отрицательных чисел:", negativeCount)
	fmt.Println("Количество нулей:", zeroCount)
}
