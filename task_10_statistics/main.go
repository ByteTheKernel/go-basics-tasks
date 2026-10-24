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

	var sum int
	var min, max int
	var evenCount, oddCount int

	for i := 0; i < n; i++ {
		var number int

		fmt.Printf("Введите число %d: ", i+1)
		if _, err := fmt.Scan(&number); err != nil {
			fmt.Println("Ошибка: необходимо ввести целое число")
			return
		}

		sum += number

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

		if number%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	average := float64(sum) / float64(n)

	fmt.Println("\nСтатистика последовательности:")
	fmt.Println("Сумма чисел:", sum)
	fmt.Printf("Среднее значение: %.2f\n", average)
	fmt.Println("Минимальное значение:", min)
	fmt.Println("Максимальное значение:", max)
	fmt.Println("Количество чётных чисел:", evenCount)
	fmt.Println("Количество нечётных чисел:", oddCount)
}
