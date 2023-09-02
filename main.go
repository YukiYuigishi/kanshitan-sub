package main

import (
	"encoding/json"
	"log"
	"net"
	//	"net/http"
	"time"

	"fmt"
	"github.com/mackerelio/go-osstat/uptime"
)

// ref https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go/37382208#37382208
func GetOutboundIP() (net.IP, error) {

	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {

		return nil, err

	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

type Status int

const (
	Up Status = iota
	Down
	None
)

type MachineData struct {
	Name   string `json:"name"`
	net.IP `json:"ip"`
	UpTime time.Duration `json:"uptime"`
	Status Status        `json:"status"`
}

func main() {
	ipaddr, err := GetOutboundIP()
	if err != nil {
		log.Fatal(err)
	}

	upTime, err := uptime.Get()
	if err != nil {
		log.Fatal(err)
	}

	data := MachineData{
		Name:   "CCD050",
		IP:     ipaddr,
		UpTime: upTime,
		Status: None,
	}
	fmt.Println("name", data.Name)
	fmt.Println("IP", data.IP)
	fmt.Println("uptime", data.UpTime)
	fmt.Println("Status", data.Status)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	test := MachineData{Name: "test", IP: ipaddr, UpTime: upTime, Status: None}
	jsonTest, _ := json.Marshal(test)
	fmt.Printf("%s\n", jsonTest)
	fmt.Println(len(jsonTest))
	fmt.Printf("%s\n", jsonData)

}
