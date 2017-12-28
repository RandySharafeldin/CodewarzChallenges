package main

import (
    
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "math"
)

func sumBetween(num1 float64, num2 float64) float64 {
    
    min := math.Min(float64(num1), float64(num2))
    max := math.Max(float64(num1), float64(num2))
    
    sum := 0.0
    for ; min <= max; min++ {
        sum += min
    }
    
    return sum
}


func main() {
    
    file, _ := os.Open(os.Args[1])
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        nums := strings.Split(scanner.Text(), " ")
        n1, _ := strconv.ParseFloat(nums[0], 10)
        n2, _ := strconv.ParseFloat(nums[1], 10)
        fmt.Println(sumBetween(n1, n2))
    }
}