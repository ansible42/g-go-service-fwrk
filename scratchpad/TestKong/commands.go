// nolint
package main

import (
	"fmt"
)

type TestCmd struct {
	Port    int32  `help:"You must provide a port number."`
	Address string `help:"You must provide a IP address"`
}

func (a *TestCmd) Run(globals *Globals) error {
	fmt.Println("This is a test")
	return nil
}
