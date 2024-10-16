package main

func main() {
    player := &Player{Name: "SecretBoy2004"}
    enemy := &Enemy{Type: "Упырь"}
    treasure := &Treasure{GoldAmount: 1000}

    achievement := &Achievement{}

    player.Accept(achievement)
    enemy.Accept(achievement)
    treasure.Accept(achievement)
}

