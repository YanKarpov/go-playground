package main

import (
    "flag"
    "fmt"
)

func main() {
    example := flag.String("example", "hello", "Choose example: hello, variables, loops")
    flag.Parse()

    switch *example {
    case "hello":
        helloWorld()   
    case "variables":
        variables()    
    case "loops":
        loops() 
    default:
        fmt.Println("Не пон")
    }
}
