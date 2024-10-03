package main

import (
	"fmt"
	"time"
)

type Payment interface {
	GetCreatedAt() int64
	GetAmount() int
	FormatDate() string
}

type TBankPayment struct {
	Id       string `json:"id"`
	Datetime string `json:"datetime"`
	Amount   int    `json:"amount"`
}

func (p TBankPayment) GetCreatedAt() int64 {
	t, err := time.Parse("2006-01-02T15:04:05", p.Datetime)
	if err != nil {
		fmt.Println("Ошибка парсинга даты:", err)
		return 0
	}
	return t.Unix()
}

func (p TBankPayment) GetAmount() int {
	return p.Amount
}

func (p TBankPayment) FormatDate() string {
	t, err := time.Parse("2006-01-02T15:04:05", p.Datetime)
	if err != nil {
		fmt.Println("Ошибка парсинга даты:", err)
		return "Неизвестная дата"
	}
	return t.Format("02.01.2006")
}

type SberPayment struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"timestamp"`
	Sum       int    `json:"sum"`
}

func (p SberPayment) GetCreatedAt() int64 {
	return p.CreatedAt
}

func (p SberPayment) GetAmount() int {
	return p.Sum
}

func (p SberPayment) FormatDate() string {
	t := time.Unix(p.CreatedAt, 0)
	return t.Format("02.01.2006")
}
