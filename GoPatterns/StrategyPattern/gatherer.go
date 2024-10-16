package main

import "fmt"

type Gatherer struct{}

func (g *Gatherer) Execute() {
	fmt.Println("Собиратель: Я соберу плоды и травы!")
}
