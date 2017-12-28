package main

import (

    "os"
    "fmt"
    "bufio"
)

func checkReverse(s []rune) bool {
    word := s
    check := true
    for i := 0; i < len(s) && check; i++ {
        check = word[len(s) - i - 1] == s[i]
    }
    
    return check
    
} 

func main() {
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        
        s := []rune(scanner.Text())
        if checkReverse(s) {
            fmt.Println("True")
        } else {fmt.Println("False")}
    }
}