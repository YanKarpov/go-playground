package main

type AggressiveState struct{}

func (a *AggressiveState) Handle() string {
	return "Агрессивный: Хищник атакует травоядное с яростью!"
}
