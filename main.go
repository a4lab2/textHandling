package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {

	f, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer f.Close()
	d, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurservers{}
	err = xml.Unmarshal(d, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}
