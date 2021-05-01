package main

import (
	"fmt"
	"net"
)

func GetComms() net.Conn {
	CONNECT := "127.0.0.1:5555"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return c
}

func SendCommand(cmd []byte) {
	c := GetComms()
	fmt.Println(string(cmd))
	c.Write(cmd)
	c.Close()
}

type CmdMsg struct {
	msg  string `json:"msg"`  //This is the name of the actor to be called.
	data string `json:"data"` //This is the private data of the called actor
}

func MakeCmd(Actor string, data string) {
	cmd := fmt.Sprintf("{\"msg\":\"%s\",\"data\":%s}\r\n", Actor, data)
	SendCommand([]byte(cmd))
}
