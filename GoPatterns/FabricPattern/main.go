package main

import (
	"fmt"
)

func main() {
	var mafiaFactory CharacterFactory = &MafiaFactory{}
	var villagerFactory CharacterFactory = &VillagerFactory{}

	mafia := mafiaFactory.CreateCharacter()
	villager := villagerFactory.CreateCharacter()

	fmt.Println(mafia.Role())
	fmt.Println(villager.Role())
}
