package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "math"
)

func countDivs(toDiv float64, max float64) {
    if max < 0 {fmt.Println(0)}
    for i := 1.0; i <= math.Abs(max); i++ {
        if math.Mod(i, toDiv) == 0 {
            if max < 0 {
            fmt.Println(-i)
            } else {
                fmt.Println(i)
            }
        }
    }
    
}

func main() {
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        nums := strings.Split(scanner.Text(), " ")
        num1, _ := strconv.ParseFloat(nums[0], 64)
        num2, _ := strconv.ParseFloat(nums[1], 64)
        countDivs(num1, num2)
        print("\n")
    }
    
}