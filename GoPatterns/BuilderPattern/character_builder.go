package main

type CharacterBuilder struct {
    character *Character
}

func NewCharacterBuilder() *CharacterBuilder {
    return &CharacterBuilder{character: &Character{}}
}

func (b *CharacterBuilder) SetName(name string) *CharacterBuilder {
    b.character.Name = name
    return b
}

func (b *CharacterBuilder) SetStrength(strength int) *CharacterBuilder {
    b.character.Strength = strength
    return b
}

func (b *CharacterBuilder) SetAgility(agility int) *CharacterBuilder {
    b.character.Agility = agility
    return b
}

func (b *CharacterBuilder) SetIntelligence(intelligence int) *CharacterBuilder {
    b.character.Intelligence = intelligence
    return b
}

func (b *CharacterBuilder) Build() *Character {
    return b.character
}
