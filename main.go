package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":10002")
	if err != nil {
		panic(err)
	}

	connection, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	buffer := make([]byte, 1024)
	for {
		count, sendAddr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("received: ", string(buffer[0:count]), " from: ", sendAddr)
		}
	}
}
