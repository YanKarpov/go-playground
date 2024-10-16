package main

type CharacterFactory interface {
	CreateCharacter() Character
}

type MafiaFactory struct{}

func (m *MafiaFactory) CreateCharacter() Character {
	return &Mafia{}
}

type VillagerFactory struct{}

func (v *VillagerFactory) CreateCharacter() Character {
	return &Villager{}
}
