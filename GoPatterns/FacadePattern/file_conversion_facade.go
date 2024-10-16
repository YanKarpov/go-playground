package main

type FileConversionFacade struct {
    reader    *FileReader
    converter *FileConverter
    saver     *FileSaver
}

func NewFileConversionFacade() *FileConversionFacade {
    return &FileConversionFacade{
        reader:    &FileReader{},
        converter: &FileConverter{},
        saver:     &FileSaver{},
    }
}

func (fcf *FileConversionFacade) ConvertFile(inputPath, outputPath, fromFormat, toFormat string) {
    content := fcf.reader.Read(inputPath)
    convertedContent := fcf.converter.Convert(content, fromFormat, toFormat)
    fcf.saver.Save(outputPath, convertedContent)
}
