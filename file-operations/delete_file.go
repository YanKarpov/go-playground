package main

import (
    "fmt"
    "os"
)

func deleteFile() {
    err := os.Remove("example.txt")
    if err != nil {
        fmt.Println("Ошибка при удалении файла:", err)
        return
    }
    fmt.Println("Файл успешно удален.")
}
