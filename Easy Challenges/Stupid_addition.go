package main

import (

    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        
        nums := strings.Split(scanner.Text(), " ")
        n1, err1 := strconv.ParseFloat(nums[0], 10)
        if err1 == nil && len(nums) == 2 {
            n2, err2 := strconv.ParseFloat(nums[1], 10)
            if err2 == nil {fmt.Printf("%.10v\n", n1 + n2)}
        }
    }
}