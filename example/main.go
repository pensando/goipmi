package main

import (
	"fmt"
	"os"
	"strconv"

	ipmi "github.com/vmware/goipmi"
)

func main() {
	conn := &ipmi.Connection{
		Hostname:  os.Args[1],
		Username:  os.Args[2],
		Password:  os.Args[3],
		Interface: "lanplus",
		Path:      "ipmitool",
	}

	client, err := ipmi.NewClient(conn)
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 4 {
		ciphersuite, err := strconv.Atoi(os.Args[4])
		if err != nil {
			panic(err)
		}
		client.SetCiphersuite(ciphersuite)
	}

	if err := client.Open(); err != nil {
		panic(err)
	}

	r := &ipmi.Request{
		NetworkFunction: ipmi.NetworkFunctionChassis,
		Command:         ipmi.CommandChassisStatus,
		Data:            &ipmi.ChassisStatusRequest{},
	}
	resp := &ipmi.ChassisStatusResponse{}
	err = client.Send(r, resp)
	if err != nil {
		panic(err)
	}
	if resp.IsSystemPowerOn() {
		fmt.Println("Power is on")
	} else {
		fmt.Println("Power is off")
	}
}
