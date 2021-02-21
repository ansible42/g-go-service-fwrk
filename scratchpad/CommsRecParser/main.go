package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/tidwall/gjson"
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
		fmt.Println(msg.ExitMessage)
	}
	os.Exit(int(msg.ExitCode))
}

type PrintLineMsg struct {
	LineToPrint string `json:"Message"`
	Color       struct {
		B uint8 `json:"b"`
		G uint8 `json:"g"`
		R uint8 `json:"r"`
	} `json:"color"`
}

func DoPrintLine(data string) {
	var msg PrintLineMsg
	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		fmt.Println(err)
	}
	color.RGB(msg.Color.R, msg.Color.G, msg.Color.B).Println(msg.LineToPrint)

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
	const ExitMsg = `{"state": "Exit", "data":{ "ExitCode": 42,"ExitMessage": "How many roads must a man walk down"}}`
	const PrintLineRed = `{"state":"PrintLine","data":{"Message":"You should print this line", "color":{"r":255,"b":0,"g":0}}}`
	const PrintLineBlue = `{"state":"PrintLine","data":{"Message":"You should print this line", "color":{"r":0,"b":255,"g":0}}}`
	const PrintLineGreen = `{"state":"PrintLine","data":{"Message":"You should print this line", "color":{"r":0,"b":0,"g":255}}}`

	GetMsgType(PrintLineRed)
	GetMsgType(PrintLineGreen)
	GetMsgType(PrintLineBlue)
	GetMsgType(ExitMsg)

	return
}
