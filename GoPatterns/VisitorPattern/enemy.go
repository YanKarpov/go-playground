package main

type Enemy struct {
    Type string
}

func (e *Enemy) Accept(visitor AchievementVisitor) {
    visitor.VisitEnemy(e)
}