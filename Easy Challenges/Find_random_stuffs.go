package main

import ( 

    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "math"
)


func likeFizzBuzz(n1 float64, n2 float64) {
    
    min := math.Min(n1, n2)
    max := math.Max(n1, n2)
    
    for math.Mod(min, 7) != 0 && min <= max {
        min++
        if math.Mod(min, 5) == 0 {min++}
    }
    fmt.Print(min)
    min++
    for ; min <= max; min++ {
        if math.Mod(min, 5) != 0 && math.Mod(min, 7) == 0 {
            fmt.Printf(",%v", min)
        }
    }
    
    print("\n")
}


func main() {
    
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        
        nums := strings.Split(scanner.Text(), " ")
        n1, _ := strconv.ParseFloat(nums[0], 10)
        n2, _ := strconv.ParseFloat(nums[1], 10)
        likeFizzBuzz(n1, n2)
    }
    
    
}