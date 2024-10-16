package main

import "fmt"

type AttackCommand struct {
    Target string
}

func (a *AttackCommand) Execute() {
    fmt.Println("Атака цели:", a.Target)
}
