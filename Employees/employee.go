package Employees

import (
	"fmt"
	"strconv"
)

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("FullTimeEmployee [ID: %d, Name: %s, Salary: %d]", f.ID, f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("PartTimeEmployee [ID: %d, Name: %s, HourlyRate: %d, HoursWorked: %.2f]",
		p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

type Company struct {
	employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{
		employees: make(map[string]Employee),
	}
}

func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.employees[strconv.FormatUint(e.ID, 10)] = e
	case PartTimeEmployee:
		c.employees[strconv.FormatUint(e.ID, 10)] = e
	default:
		fmt.Println("Неизвестный тип сотрудника")
		return
	}
	fmt.Println("Сотрудник добавлен:", emp.GetDetails())
}
func (c *Company) ListEmployees() {
	fmt.Println("Список")
	for _, emp := range c.employees {
		fmt.Println(emp.GetDetails())
	}
	fmt.Println()
}
