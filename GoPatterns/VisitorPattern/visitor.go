package main

import "fmt"

type AchievementVisitor interface {
    VisitPlayer(player *Player)
    VisitEnemy(enemy *Enemy)
    VisitTreasure(treasure *Treasure)
}

type Achievement struct{}

func (a *Achievement) VisitPlayer(player *Player) {
    fmt.Println("Достижение: победил другого игрока", player.Name)
}

func (a *Achievement) VisitEnemy(enemy *Enemy) {
    fmt.Println("Достижение: победил первого врага", enemy.Type)
}

func (a *Achievement) VisitTreasure(treasure *Treasure) {
    fmt.Println("Достижение: найдено сокровище с количеством золота:", treasure.GoldAmount)
}