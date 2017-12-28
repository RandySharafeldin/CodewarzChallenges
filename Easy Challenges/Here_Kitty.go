package main

import (
    "os"
    "fmt"
    "io/ioutil"
)

func main() {
    dat, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
    panic(err)
    }
    fmt.Print(string(dat))
}