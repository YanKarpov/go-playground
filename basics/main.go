package main

import (
    "flag"
    "fmt"
)

func main() {
    example := flag.String("example", "hello", "Choose example: hello, variables, loops, structs")
    flag.Parse()

    switch *example {
    case "hello":
        helloWorld()   
    case "variables":
        variables()    
    case "loops":
        loops() 
    case "structs":
        structs() 
    default:
        fmt.Println("Не пон")
    }
}