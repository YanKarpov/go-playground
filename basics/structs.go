package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Метод для структуры Person, который выводит информацию
func (p Person) Info() {
    fmt.Printf("Имя: %s, Возраст: %d\n", p.Name, p.Age)
}

// Метод для изменения возраста
func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Printf("%s теперь %d лет!\n", p.Name, p.Age)
}

// Метод для изменения имени
func (p *Person) ChangeName(newName string) {
    p.Name = newName
    fmt.Printf("Имя изменено на: %s\n", p.Name)
}

// Функция для работы со структурой Person
func structs() {
    p := Person{Name: "Ян", Age: 19}
    p.Info()
    p.HaveBirthday()
    p.ChangeName("Антон")
    p.Info() 
}
