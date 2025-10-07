package main

import (
	"fmt"
)

type BankAccount struct {
	owner   string
	number  string
	balance float64
}

// конструктор - владелец, номер счета, баланс
func NewBankAccount(owner string, number string, balance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		number:  number,
		balance: balance,
	}
}

// Геттер имени владельца
func (acc BankAccount) GetOwner() string {
	return acc.owner
}

// Геттер имени владельца
func (acc BankAccount) GetBalance() float64 {
	return acc.balance
}

// валидация депозита и депозит dep на счёт acc (возврат статуса)
func (acc *BankAccount) DepValid(dep float64) bool {
	if dep > 0 {
		acc.balance += dep
		return true
	} else {
		return false
	}
}

// метод Deposit(x) добавляет dep средств на баланс счёта acc
func Deposit(acc *BankAccount, dep float64) {
	fmt.Printf("Зачисление депозита %.00fр на счет %s\n", dep, acc.number)
	if acc.DepValid(dep) {
		fmt.Println("Операция выполнена!")
	} else {
		fmt.Println("Операция не выполнена, зачисление должно быть положительным!")
	}
}

// валидация снятия и снятие суммы with со счета acc (возврат статуса)
func (acc *BankAccount) WithValid(with float64) bool {
	if with <= acc.balance && with > 0 {
		acc.balance -= with
		return true
	} else {
		return false
	}
}

// функция для снятия суммы with со счёта acc
func Withdraw(acc *BankAccount, with float64) {
	fmt.Printf("Снятие %.00fр со счета %s\n", with, acc.number)
	if acc.WithValid(with) {
		fmt.Println("Операция выполнена!")
	} else {
		fmt.Println("Операция не выполнена, попытка снять отрицательную сумму, либо недостаточно средств на счете!")
	}
}

func main() {
	// новый счёт на имя Mary
	account := NewBankAccount("Mary", "0000 0000 0000 0000", 0)

	// успешное зачисление 500 и снятие 200
	Deposit(account, 500)
	Withdraw(account, 250)

	// ошибка отрицательного депозита и недостаточности средств
	Deposit(account, -100)
	Withdraw(account, 1000)

	// вывод баланса Mary с помощью геттеров
	fmt.Println(account.GetOwner(), " balance: ", account.GetBalance())
}
