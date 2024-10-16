package main

import "fmt"

type FileReader struct{}

func (fr *FileReader) Read(filePath string) string {
    fmt.Println("Чтение файла:", filePath)
    return "file_content"
}
