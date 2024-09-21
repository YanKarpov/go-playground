package main

import (
    "flag"
    "fmt"
)

func main() {
    operation := flag.String("operation", "add", "Choose operation: add, copy, delete")
    flag.Parse()

    switch *operation {
    case "add":
        addToFile()
    case "copy":
        copyFile()
    case "delete":
        deleteFile()
    default:
        fmt.Println("Неизвестная операция.")
    }
}
