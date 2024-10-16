package main

import "fmt"

type Character struct {
    Name       string
    Strength   int
    Agility    int
    Intelligence int
}

func (c *Character) ShowInfo() {
    fmt.Printf("Персонаж: %s\nСила: %d\nЛовкость: %d\nИнтеллект: %d\n", 
        c.Name, c.Strength, c.Agility, c.Intelligence)
}
