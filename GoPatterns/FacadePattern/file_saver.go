package main

import "fmt"

type FileSaver struct{}

func (fs *FileSaver) Save(filePath, content string) {
    fmt.Println("Сохранение файла по пути:", filePath)
}
