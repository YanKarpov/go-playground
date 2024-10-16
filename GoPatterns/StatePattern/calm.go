package main

type CalmState struct{}

func (c *CalmState) Handle() string {
	return "Спокойный: Хищник сохраняет спокойствие и наблюдает за врагом."
}
