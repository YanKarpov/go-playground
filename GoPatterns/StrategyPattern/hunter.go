package main

import "fmt"

type Hunter struct{}

func (h *Hunter) Execute() {
	fmt.Println("Охотник - Я буду охотиться на дичь!")
}
