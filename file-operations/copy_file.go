package main

import (
    "fmt"
    "io"
    "os"
)

func copyFile() {
    srcFile, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Ошибка при открытии исходного файла:", err)
        return
    }
    defer srcFile.Close()

    dstFile, err := os.Create("copy_example.txt")
    if err != nil {
        fmt.Println("Ошибка при создании файла копии:", err)
        return
    }
    defer dstFile.Close()

    if _, err := io.Copy(dstFile, srcFile); err != nil {
        fmt.Println("Ошибка при копировании файла:", err)
        return
    }
    fmt.Println("Файл успешно скопирован.")
}
