// nolint
package main

import (
	"encoding/json"
	"github.com/gookit/color"
)

type CountCmd struct {
	CountI int32 `arg help:"Provide a count of the number to return."`
}

func (a *CountCmd) Run(globals *Globals) error {
	CommandType := "TestProcesses.lvlib:Counter.lvclass"
	j, e := json.Marshal(a)
	if e != nil {
		color.Red.Println(e)
	}
	HandleComms(CommandType, j)
	return nil
}
