package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/tidwall/gjson"
	"net"
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

func GetMsgType(JSONBall string) {
	state := gjson.Get(JSONBall, "state") // state is a known and required fiend
	data := gjson.Get(JSONBall, "data")   // data can be anything
	//fmt.Println("Msg State :: ", state)
	//fmt.Println("Msg Data :: ", data.String())
	switch state.String() {
	case "Exit":
		DoExitMsg(data.String())
	case "PrintLine":
		DoPrintLine(data.String())
	}

}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	fmt.Println("Connecting to ", CONNECT)
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Println("Waiting for message")
		LatestMsg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			color.Red.Println(err)
		}
		GetMsgType(LatestMsg)
	}
}
