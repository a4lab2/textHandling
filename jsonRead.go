package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp   string
}
type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	str := `{"servers":[{"serverName":"zone_A","serverIP":"127.0.0.1"},{"serverName":"zone_B","serverIP":"
127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
