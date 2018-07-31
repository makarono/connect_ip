package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	hosts := [4]string{"10.22.25.26", "192.168.33.1", "10.11.11.11", "12.12.12.10"}

	portNum := "22"
	seconds := 10
	timeOut := time.Duration(seconds) * time.Second

	i := 0
	for _, ip := range hosts {
		//fmt.Println(ip)
		conn, err := net.DialTimeout("tcp", ip+":"+portNum, timeOut)
		i++
		//fmt.Println("counter:", i)

		if err != nil {
			fmt.Println("CRITICAL:", err)
			os.Exit(1)
		}
		//fmt.Printf("Connection established between %s and localhost with time out of %d seconds.\n", ip, int64(seconds))

		fmt.Printf("Connected to: %s \n", conn.RemoteAddr().String())
		//			fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

	}

}
