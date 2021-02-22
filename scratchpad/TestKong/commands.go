// nolint
package main

import (
	"fmt"
	"path"
)

type TestCmd struct {
	Port    int32  `arg help:"You must provide a port number."`
	Address string `arg help:"You must provide a IP address"`
}

func (a *TestCmd) Run(globals *Globals) error {
	fmt.Println("This is a test")
	return nil
}
