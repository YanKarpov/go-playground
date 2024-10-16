package main

import "fmt"

type FileConverter struct{}

func (fc *FileConverter) Convert(content, fromFormat, toFormat string) string {
    fmt.Printf("Конвертация из %s в %s...\n", fromFormat, toFormat)
    return "converted_content"
}
