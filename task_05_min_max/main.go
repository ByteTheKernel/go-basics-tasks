package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число")
		return
	}

	if n <= 0 {
		fmt.Println("Ошибка: N должно быть положительным")
		return
	}

	var number, max, min int

	for i := 0; i < n; i++ {
		fmt.Printf("Введите число %d: ", i+1)

		if _, err := fmt.Scan(&number); err != nil {
			fmt.Println("Ошибка: необходимо ввести целое число")
			return
		}

		if i == 0 {
			min = number
			max = number
		} else {
			if number < min {
				min = number
			}

			if number > max {
				max = number
			}
		}
	}

	fmt.Println("Минимальное значение:", min)
	fmt.Println("Максимальное значение:", max)
}
