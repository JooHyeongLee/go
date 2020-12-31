package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:5555")
	if err != nil {
		log.Panic(err)
	}

	socket, err := net.DialUDP("udp4", nil, ServerAddr)

	if err != nil {
		log.Panic(err)
	}

	// Server에서 값을 받는 고루틴
	go func() {
		data := make([]byte, 4096)
		for {
			n, addr, err := socket.ReadFromUDP(data)
			// n, err := socket.Read(data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			log.Printf("Server send : %v, ip : %v", string(data[:n]), addr)
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	// 사용자 입력
	for {
		var s string
		fmt.Scanln(&s)
		socket.Write([]byte(s))
		time.Sleep(time.Duration(3) * time.Second)
	}

	//defer socket.Close()

	//data := "hello"
	//socket.Write([]byte(data))
	//log.Println("C-OUT] " + data)
}
