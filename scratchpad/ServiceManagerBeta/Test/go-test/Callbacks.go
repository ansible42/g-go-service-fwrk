package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"os"
	_ "strings"
)

type ExitMsg struct {
	ExitCode    int64  `json:"ExitCode"`
	ExitMessage string `json:"ExitMessage"`
}

func DoExitMsg(data string) {
	var msg ExitMsg
	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		fmt.Println(err)
	}
	if msg.ExitCode != 0 {
		color.Red.Print(msg.ExitMessage, true)
	}
	os.Exit(int(msg.ExitCode))
}

type PrintLineMsg struct {
	LineToPrint string `json:"LineToPrint"`
	TxtColor    struct {
		B uint8 `json:"b"`
		G uint8 `json:"g"`
		R uint8 `json:"r"`
	} `json:"TxtColor"`
}

func DoPrintLine(data string) {
	var msg PrintLineMsg
	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		fmt.Println(err)
	}
	color.RGB(msg.TxtColor.R, msg.TxtColor.G, msg.TxtColor.B).Println(msg.LineToPrint)

}
