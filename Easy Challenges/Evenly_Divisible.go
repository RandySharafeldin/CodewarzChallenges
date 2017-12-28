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
    i := 1.0
    if toDiv > 1.0 {
        i = math.Floor(toDiv)
    }
    
    if math.Mod(i, toDiv) == 0 {
        fmt.Print(i)
        i++
    }
    
    for ; i <= math.Abs(max); i++ {
        if math.Mod(i, toDiv) == 0 {
            print("\n")
            if max < 0 {
            fmt.Print(-i)
            } else {
                fmt.Print(i)
            }
        }

    }
    
    if math.Mod(max, toDiv) == 0 {
        print("\n")
    }
    
}

func main() {
    
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s := scanner.Text()
        nums := strings.Split(s, " ")
        num1, _ := strconv.ParseFloat(nums[0], 64)
        num2, _ := strconv.ParseFloat(nums[1], 64)
        countDivs(num1, num2)
        nums = strings.Split(s, " ")
        print("\n")
        }

    }
    