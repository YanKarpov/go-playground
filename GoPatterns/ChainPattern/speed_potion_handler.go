package main

import "fmt"

type SpeedPotionHandler struct {
    BaseHandler
}

func (h *SpeedPotionHandler) UseItem(item string) {
    if item == "speed_potion" {
        fmt.Println("Использовано зелье скорости: +30 к скорости")
    } else {
        fmt.Println("Зелье скорости не подходит, передаем дальше...")
        h.BaseHandler.UseItem(item)
    }
}
