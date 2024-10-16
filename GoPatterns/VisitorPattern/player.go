package main

type Player struct {
    Name string
}

func (p *Player) Accept(visitor AchievementVisitor) {
    visitor.VisitPlayer(p)
}