package main

import "fmt"

type DefendCommand struct{}

func (d *DefendCommand) Execute() {
    fmt.Println("Защита активирована")
}
