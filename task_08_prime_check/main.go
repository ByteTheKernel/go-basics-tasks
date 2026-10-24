package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int

	fmt.Print("Введите целое число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число")
		return
	}

	if isPrime(n) {
		fmt.Printf("Число %d является простым\n", n)
	} else {
		fmt.Printf("Число %d не является простым\n", n)
	}
}
