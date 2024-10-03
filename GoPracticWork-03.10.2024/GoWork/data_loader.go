package main

import (
	"encoding/json"
	"io"
	"os"
)

type RawPayment struct {
	Id        string `json:"id"`
	Datetime  string `json:"datetime,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Amount    int    `json:"amount,omitempty"`
	Sum       int    `json:"sum,omitempty"`
	Bank      string `json:"bank"`
}

func ParsePayments(filePath string) ([]Payment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var rawPayments []RawPayment
	if err := json.Unmarshal(data, &rawPayments); err != nil {
		return nil, err
	}

	payments := make([]Payment, 0, len(rawPayments))

	for _, raw := range rawPayments {
		switch raw.Bank {
		case "Sber":
			payments = append(payments, SberPayment{
				Id:        raw.Id,
				CreatedAt: raw.Timestamp,
				Sum:       raw.Sum,
			})
		case "T-Bank":
			payments = append(payments, TBankPayment{
				Id:       raw.Id,
				Datetime: raw.Datetime,
				Amount:   raw.Amount,
			})
		}
	}

	return payments, nil
}
