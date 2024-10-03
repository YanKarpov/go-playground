package main

import (
	"fmt"
)

func main() {
	payments, err := ParsePayments("payments.json")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	for _, payment := range payments {
		fmt.Printf("Дата создания: %s, Сумма: %d\n", payment.FormatDate(), payment.GetAmount())
	}
}





