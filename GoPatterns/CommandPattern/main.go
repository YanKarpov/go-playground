package main


func main() {
    attack := AttackCommand{Target: "Дракон"}
    defend := DefendCommand{}
    useItem := UseItemCommand{Item: "Зелье здоровья"}

    attack.Execute()
    defend.Execute()
    useItem.Execute()
}
