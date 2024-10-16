package main

import (
    "fmt"
)

func main() {
    // Создаем персонажа с начальным спокойным состоянием
    character := NewCharacter(&CalmState{})
    fmt.Println(character.Act())

    // Переключаемся на агрессивное состояние
    character.SetState(&AggressiveState{})
    fmt.Println(character.Act())

    // Переключаемся на защитное состояние
    character.SetState(&DefensiveState{})
    fmt.Println(character.Act())
}
