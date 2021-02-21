package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"

	"github.com/alecthomas/kong"
)

type t_msg struct {
	state string
	data  string
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Println("Waiting for message")
		var LatestMsg t_msg
		message, _ := bufio.NewReader(c).ReadString('\n')
		println("RAW MSG:: ", message)
		LatestMsg.state = message[0:4]
		LatestMsg.data = message[5:]
		PrintMsgRec(LatestMsg)
		ParseMsgRec(LatestMsg)
	}
}

func PrintMsgRec(Message t_msg) {
	fmt.Println("Message State:: ", Message.state)
	fmt.Println("Message DATA:: ", Message.data)
}

func ParseMsgRec(Message t_msg) {
	RunMsg := func(ctx context.Context, t_msg.Message.state) {

	}
}
