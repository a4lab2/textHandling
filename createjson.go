package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	s.Servers = append(s.Servers, Server{ServerName: "zone_A", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "zone_B", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)

	}
	fmt.Println(string(b))
}
