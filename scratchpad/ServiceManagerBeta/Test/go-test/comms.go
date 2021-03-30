package main

import (
	"bufio"
	//"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/tidwall/gjson"
	"net"
	_ "strings"
)

func HandleComms(Type string, Data []byte) {
	CONNECT := "127.0.0.1:5559"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	packet := FormatCommand(Type, Data)
	c.Write(packet)
	for {
		fmt.Println("Waiting for message")
		LatestMsg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			color.Red.Println(err)
		}
		GetMsgType(LatestMsg)

	}
}

func FormatCommand(Type string, Data []byte) []byte {
	j := fmt.Sprintf("{\"msg\":\"%s\",\"data\":%s}\r\n", Type, string(Data))

	println(string(j))
	return []byte(j)
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
