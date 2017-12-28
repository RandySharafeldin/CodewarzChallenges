package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func main() {
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    reader := bufio.NewReader(file)
    
    s, _, _ := reader.ReadLine()
    
    words := strings.Split(string(s), ". ")
    
    for _, w := range words {
        fmt.Printf("%v ", w[0:strings.IndexAny(w, " ")])
    }
    
}