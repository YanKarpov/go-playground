package main

import "fmt"

type HealthPotionHandler struct {
    BaseHandler
}

func (h *HealthPotionHandler) UseItem(item string) {
    if item == "health_potion" {
        fmt.Println("Использовано зелье здоровья: +50 HP")
    } else {
        fmt.Println("Зелье здоровья не подходит, передаем дальше...")
        h.BaseHandler.UseItem(item)
    }
}
