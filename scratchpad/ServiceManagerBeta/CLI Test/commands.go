// nolint
package main

import (
	"encoding/json"
	"fmt"
)

type TestCmd struct {
	Port    int32  `arg help:"You must provide a port number."`
	Address string `arg help:"You must provide a IP address"`
}

func (a *TestCmd) Run(globals *Globals) error {
	fmt.Println("This is a test")
	return nil
}

type CountdownCmd struct {
	Count int32 `arg required help:The count number to return.`
}

func (a *CountdownCmd) Run(globals *Globals) error {
	// Marshall the private data for the class to be called
	privateData, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//the actor name is hard coded to its client class
	MakeCmd("TestProcesses.lvlib:Counter.lvclass", string(privateData))
	return nil
}
