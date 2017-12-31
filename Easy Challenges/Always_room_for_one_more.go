package main

import (
	"os"
	"fmt"
	"bufio"
	"net"
	"strconv"
	//"io/ioutil"
)

func main() {
	addr := fmt.Sprintf("%v:%v", os.Args[1], os.Args[2])
	tcpAddr , err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		print("connection failed")
	}
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//writer := bufio.NewWriter(conn)

	for {
		//reader.ReadLine()
		s, _, _ := reader.ReadLine()
		//reader.ReadLine()
		//fmt.Println(string(s))
		if len(s) > 13 {
			n, err := strconv.ParseInt(string(s[13:len(s)-1]), 10, 64)
			if err != nil {
				print("broke here")
				break
			}
			fmt.Println(n + 1)
			_, err = conn.Write([]byte(fmt.Sprintf("%s", n + 1)))
			//s, _, _ = reader.ReadLine()
			fmt.Print(string(s))

			if err != nil {
				print("kill moe")
			}
		}


	}
}
