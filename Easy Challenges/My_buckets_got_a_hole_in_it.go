package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strconv"
    "strings"
)

func countTens(f *os.File, nums []int) int{
    
    max := 0
    
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        curr, _ := strconv.Atoi(scanner.Text())
        n := curr / 10
        nums[n]++
        if n > max {
            max = n
        }
    }
    
    return max
}

func main() {
    
    var nums [100]int
    
    file, err := os.Open(os.Args[1])
    if err != nil {log.Fatal(err)}
    
    defer file.Close()
    
    n := countTens(file, nums[:])
    fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(nums[:n+1]), "[]"), " ", "", -1)) 
}