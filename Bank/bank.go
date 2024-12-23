package Bank

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма депозита не может быть меньше нуля")
		return
	}
	b.Balance += amount
	fmt.Printf("Депозит %.2f успешно выполнен. Текущий баланс: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма снятия не может быть меньше или равна 0")
		return
	}
	if b.Balance < amount {
		fmt.Println("Недостаточно средств для снятия.")
		return
	}
	b.Balance -= amount
	fmt.Printf("Снятие %.2f успешно выполнено. Текущий баланс: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Текущий баланс: %.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, t := range transactions {
		if t > 0 {
			account.Deposit(t)
		} else {
			account.Withdraw(-t)
		}
	}
}
