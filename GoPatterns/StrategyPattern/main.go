package main

import (
	"fmt"
)

func main() {
	hunter := &Hunter{}
	gatherer := &Gatherer{}

	survivor1 := &Survivor{strategy: hunter}
	survivor2 := &Survivor{strategy: gatherer}

	fmt.Println("Выживший 1:")
	survivor1.strategy.Execute()

	fmt.Println("Выживший 2:")
	survivor2.strategy.Execute()

	fmt.Println("\nВыживший 1 меняет стратегию на собирателя:")
	survivor1.SetStrategy(gatherer)
	survivor1.strategy.Execute()
}

type Survivor struct {
	strategy SurvivalStrategy
}

func (s *Survivor) SetStrategy(strategy SurvivalStrategy) {
	s.strategy = strategy
}


