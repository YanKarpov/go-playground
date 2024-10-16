package main

// Character - контекст, который изменяет своё поведение в зависимости от состояния
type Character struct {
    state State
}

// NewCharacter создает персонажа с начальным состоянием
func NewCharacter(state State) *Character {
    return &Character{state: state}
}

// SetState изменяет текущее состояние персонажа
func (c *Character) SetState(state State) {
    c.state = state
}

// Act выполняет действие в зависимости от текущего состояния
func (c *Character) Act() string {
    return c.state.Handle()
}
