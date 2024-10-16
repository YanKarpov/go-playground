package main

type Treasure struct {
    GoldAmount int
}

func (t *Treasure) Accept(visitor AchievementVisitor) {
    visitor.VisitTreasure(t)
}