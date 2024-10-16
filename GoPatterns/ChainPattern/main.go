package main

import (
    "fmt"
)

func main() {
    healthHandler := &HealthPotionHandler{}
    shieldHandler := &ShieldPotionHandler{}
    speedHandler := &SpeedPotionHandler{}

    healthHandler.SetNext(shieldHandler).SetNext(speedHandler)

    fmt.Println("Используем 'health_potion':")
    healthHandler.UseItem("health_potion")

    fmt.Println("\nИспользуем 'shield_potion':")
    healthHandler.UseItem("shield_potion")

    fmt.Println("\nИспользуем 'speed_potion':")
    healthHandler.UseItem("speed_potion")

    fmt.Println("\nИспользуем 'unknown_potion':")
    healthHandler.UseItem("unknown_potion")
}

