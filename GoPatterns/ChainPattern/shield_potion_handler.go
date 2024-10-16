package main

import "fmt"

type ShieldPotionHandler struct {
    BaseHandler
}

func (h *ShieldPotionHandler) UseItem(item string) {
    if item == "shield_potion" {
        fmt.Println("Использовано зелье защиты: +20 к защите")
    } else {
        fmt.Println("Зелье защиты не подходит, передаем дальше...")
        h.BaseHandler.UseItem(item)
    }
}
