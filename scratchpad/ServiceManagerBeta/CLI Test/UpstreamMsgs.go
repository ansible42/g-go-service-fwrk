package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"os"
)

// Each Upstream message is defined by a json struct and has an accompanying function
//

// Exit message
type ExitMsg struct {
	ExitCode    int64  `json:"ExitCode"`
	ExitMessage string `json:"ExitMessage"`
}

// Exit returns the CLI to the user with an os error
func DoExitMsg(data string) {
	var msg ExitMsg
	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		fmt.Println("Error parsing message")
		fmt.Println(err)
	}
	if msg.ExitCode != 0 {
		color.Red.Print(msg.ExitMessage, true)
	}
	os.Exit(int(msg.ExitCode))
}

// Prints a line sent by the service in the designated color
type PrintLineMsg struct {
	LineToPrint string `json:"LineToPrint"`
	TxtColor    struct {
		B uint8 `json:"b"`
		G uint8 `json:"g"`
		R uint8 `json:"r"`
	} `json:"TxtColor"`
}
