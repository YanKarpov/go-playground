package main

import "fmt"

type UseItemCommand struct {
    Item string
}

func (u *UseItemCommand) Execute() {
    fmt.Println("Используется предмет:", u.Item)
}
