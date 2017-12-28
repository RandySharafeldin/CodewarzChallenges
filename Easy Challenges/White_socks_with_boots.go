package main

import (
    
    "os"
    "net"
    "fmt"
	"io/ioutil"
	//"strings"
)

func main() {
    addr := fmt.Sprintf("%v:%v", os.Args[1], os.Args[2])
	conn, _ := net.Dial("tcp", addr)
	defer conn.Close()
	s , _ := ioutil.ReadAll(conn)
    fmt.Printf("%s", s)
}