package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"Main/Банк"
	"Main/Библиотека"
	"Main/Сотрудники"
	"Main/Шейпы"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Exercise 1: Library Demo ===")
	library := Library.NewLibrary()

	for {
		fmt.Println("Выберите действие для библиотеки (Add, Borrow, Return, List, Next):")
		if !scanner.Scan() {
			break
		}
		command := strings.TrimSpace(scanner.Text())

		if command == "Add" {
			// Добавить книгу
			fmt.Println("Введите ID, Title, Author (через запятую):")
			if !scanner.Scan() {
				break
			}
			input := scanner.Text()
			parts := strings.Split(input, ",")
			if len(parts) < 3 {
				fmt.Println("Ошибка ввода.")
				continue
			}
			id := strings.TrimSpace(parts[0])
			title := strings.TrimSpace(parts[1])
			author := strings.TrimSpace(parts[2])
			book := Library.Book{
				ID:         id,
				Title:      title,
				Author:     author,
				IsBorrowed: false,
			}
			library.AddBook(book)
		} else if command == "Borrow" {
			fmt.Println("Введите ID книги, которую хотите взять:")
			if !scanner.Scan() {
				break
			}
			id := strings.TrimSpace(scanner.Text())
			library.BorrowBook(id)
		} else if command == "Return" {
			fmt.Println("Введите ID книги, которую хотите вернуть:")
			if !scanner.Scan() {
				break
			}
			id := strings.TrimSpace(scanner.Text())
			library.ReturnBook(id)
		} else if command == "List" {
			library.ListBooks()
		} else if command == "Next" {
			// Переходим к Exercise 2
			break
		} else {
			fmt.Println("Неизвестная команда.")
		}
	}

	//----------------------------------------------------------------
	// 2. Демонстрация Exercise 2 (Shapes)
	fmt.Println("\n=== Exercise 2: Shapes Demo ===")
	shapes := []Shapes.Shape{
		Shapes.Rectangle{Length: 10, Width: 5},
		Shapes.Circle{Radius: 7},
		Shapes.Square{Length: 4},
		Shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapes {
		Shapes.PrintShapeDetails(shape)
	}

	fmt.Println("Нажмите Enter, чтобы перейти к Exercise 3...")
	scanner.Scan()

	//----------------------------------------------------------------
	// 3. Демонстрация Exercise 3 (Employee management)
	fmt.Println("\n=== Exercise 3: Employee Management Demo ===")
	company := Employees.NewCompany()

	for {
		fmt.Println("Выберите действие: AddFull, AddPart, List, Next")
		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(scanner.Text())

		if cmd == "AddFull" {
			// Добавляем FullTimeEmployee
			fmt.Println("Введите: ID, Name, Salary (через запятую):")
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			parts := strings.Split(line, ",")
			if len(parts) < 3 {
				fmt.Println("Неверный ввод!")
				continue
			}
			idStr := strings.TrimSpace(parts[0])
			name := strings.TrimSpace(parts[1])
			salaryStr := strings.TrimSpace(parts[2])

			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				fmt.Println("ID должен быть числом!")
				continue
			}
			salary, err := strconv.ParseUint(salaryStr, 10, 32)
			if err != nil {
				fmt.Println("Salary должен быть целым числом!")
				continue
			}

			fullEmp := Employees.FullTimeEmployee{
				ID:     id,
				Name:   name,
				Salary: uint32(salary),
			}
			company.AddEmployee(fullEmp)

		} else if cmd == "AddPart" {
			// Добавляем PartTimeEmployee
			fmt.Println("Введите: ID, Name, HourlyRate, HoursWorked (через запятую):")
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			parts := strings.Split(line, ",")
			if len(parts) < 4 {
				fmt.Println("Неверный ввод!")
				continue
			}
			idStr := strings.TrimSpace(parts[0])
			name := strings.TrimSpace(parts[1])
			hourlyRateStr := strings.TrimSpace(parts[2])
			hoursWorkedStr := strings.TrimSpace(parts[3])

			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				fmt.Println("ID должен быть числом!")
				continue
			}
			hourlyRate, err := strconv.ParseUint(hourlyRateStr, 10, 64)
			if err != nil {
				fmt.Println("HourlyRate должен быть целым числом!")
				continue
			}
			hoursWorked, err := strconv.ParseFloat(hoursWorkedStr, 32)
			if err != nil {
				fmt.Println("HoursWorked должен быть числом с плавающей точкой!")
				continue
			}

			partEmp := Employees.PartTimeEmployee{
				ID:          id,
				Name:        name,
				HourlyRate:  hourlyRate,
				HoursWorked: float32(hoursWorked),
			}
			company.AddEmployee(partEmp)

		} else if cmd == "List" {
			company.ListEmployees()
		} else if cmd == "Next" {
			break
		} else {
			fmt.Println("Неизвестная команда")
		}
	}

	fmt.Println("\n=== Exercise 4: Bank Account Demo ===")
	account := &Bank.BankAccount{
		AccountNumber: "123456",
		HolderName:    "Ivan Petrov",
		Balance:       0,
	}

	for {
		fmt.Println("Выберите действие: Deposit, Withdraw, Get balance, Exit")
		if !scanner.Scan() {
			break
		}
		command := strings.TrimSpace(scanner.Text())
		if command == "Deposit" {
			fmt.Println("Введите сумму депозита:")
			if !scanner.Scan() {
				break
			}
			amountStr := strings.TrimSpace(scanner.Text())
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Некорректная сумма!")
				continue
			}
			account.Deposit(amount)
		} else if command == "Withdraw" {
			fmt.Println("Введите сумму снятия:")
			if !scanner.Scan() {
				break
			}
			amountStr := strings.TrimSpace(scanner.Text())
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Некорректная сумма!")
				continue
			}
			account.Withdraw(amount)
		} else if command == "Get balance" {
			account.GetBalance()
		} else if command == "Exit" {
			fmt.Println("Программа завершена")
			break
		} else {
			fmt.Println("Неизвестная команд")
		}
	}

	fmt.Println("Демонстрация всех упражнений завер")
}
