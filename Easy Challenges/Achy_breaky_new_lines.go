package main

import (

    "os"
    "fmt"
    "bufio"
    //"strings"
)

func main() {
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    
    reader := bufio.NewReader(file)
    
    s, _, err := reader.ReadLine()
    
    for err == nil {
        
        if len(s) != 0 {
            fmt.Println(string(s))
        }
        s, _ , err = reader.ReadLine()
    }
}