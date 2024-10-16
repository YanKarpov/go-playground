package main


func main() {
    builder := NewCharacterBuilder()
    character := builder.SetName("Ведьмак").SetStrength(80).SetAgility(60).SetIntelligence(50).Build()
    character.ShowInfo()
}
