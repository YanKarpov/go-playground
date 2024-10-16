package main

type DefensiveState struct{}

func (d *DefensiveState) Handle() string {
	return "Защитный: Хищник принимает оборонительную стойку."
}
