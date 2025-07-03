package main

import (
	"encoding/json"
	"fmt"
	"kenaito-ansible-go-example/core"
)

func main() {
	testExecuteCommand()
	//testExecutePlaybook()
}
func testExecuteCommand() {
	command := core.ExecuteCommand(
		"local",
		[]string{"192.168.150.170", "192.168.205.234"},
		"command",
		"pingxxxx 127.0.0.1 -c 2",
		nil,
	)
	marshal, err := json.Marshal(command)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Printf("Execute Result: %s\n", string(marshal))
}
func testExecutePlaybook() {
	command := core.ExecutePlaybook(
		"local",
		[]string{"192.168.150.170", "192.168.205.234"},
		[]string{"site1.yml", "site2.yml", "site3.yml"},
		nil,
	)
	marshal, err := json.Marshal(command)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Printf("Execute Playbook Result: %s\n", string(marshal))
}
