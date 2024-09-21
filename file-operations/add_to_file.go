package main

import (
    "fmt"
    "os"
)

func addToFile() {
    f, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Ошибка при создании файла:", err)
        return
    }
    defer f.Close()

    _, err = f.WriteString("Привет, булочка с корицей! Что ты тут ищешь?")
    if err != nil {
        fmt.Println("Ошибка при записи в файл:", err)
        return
    }
    
    fmt.Println("Файл успешно создан и содержимое записано.")
}

