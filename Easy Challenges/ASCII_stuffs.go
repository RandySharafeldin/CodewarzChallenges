package main

import (
    "os"
    "fmt"
    "strconv"
    "bufio"
    "strings"
)

func main() {
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    data := bufio.NewScanner(file)
    data.Scan()
    words := strings.Split(data.Text(), " ")
    for _ , w := range words {
        c, _ := strconv.ParseInt(w, 16, 32)
        fmt.Printf("%c", c)
    }
}