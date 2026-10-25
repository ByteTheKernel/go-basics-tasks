# Основы Go — 10 практических заданий

[![Go CI](https://github.com/ByteTheKernel/go-basics-tasks/actions/workflows/go.yml/badge.svg)](https://github.com/ByteTheKernel/go-basics-tasks/actions/workflows/go.yml)

Учебная работа по дисциплине «Семинар наставника».

Выполнил: Орешко Владислав

## Запуск

Требуется Go 1.22 или новее. Например:

```bash
go run ./task_01_calculator
```

Замените `01_calculator` на название нужного задания от `01_calculator` до `10_statistics`. Каждая программа запрашивает данные в терминале.

## Задания

| Папка | Тема |
|---|---|
| `task_01_calculator`            | Калькулятор |
| `task_02_number_analysis`       | Анализ числа |
| `task_03_sum_average`           | Сумма и среднее |
| `task_04_multiplication_table`  | Таблица умножения |
| `task_05_min_max`               | Минимум и максимум |
| `task_06_number_counts`         | Подсчёт знаков |
| `task_07_conditional_sum`       | Сумма подходящих чисел |
| `task_08_prime_check`           | Простое число |
| `task_09_atm`                   | Мини-банкомат |
| `task_10_statistics`            | Статистика последовательности |

## Проверка

```bash
go test ./...
go vet ./...
```
