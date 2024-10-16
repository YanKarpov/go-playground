package main


func main() {
    conversionFacade := NewFileConversionFacade()
    
    inputPath := "input.txt"
    outputPath := "output.mp3"
    fromFormat := "txt"
    toFormat := "docx"

    conversionFacade.ConvertFile(inputPath, outputPath, fromFormat, toFormat)
}
