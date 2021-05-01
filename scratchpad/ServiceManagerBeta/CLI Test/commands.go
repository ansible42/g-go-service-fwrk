// nolint
package main

import (
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

	return nil
}
