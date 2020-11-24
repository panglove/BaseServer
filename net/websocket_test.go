package net

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"strings"
	"testing"
	"time"
)

var ClientInstall *websocket.Conn

func TestNet(t *testing.T) {

	url := "ws://127.0.0.1:8083/ws"

	wsCon, err := websocket.Dial(url, "", "http://127.0.0.1")

	if err != nil {
		log.Fatal("ws connect err:", err)
	}
	ClientInstall = wsCon

	go func() {

		for {
			err := websocket.Message.Send(wsCon, "hello")
			if err != nil {
				fmt.Println("send err:", err)

			} else {

				fmt.Println("send msg")
			}
			time.Sleep(time.Second * 5)

		}
	}()
	GoClientReceive()

	select {}

}
func GoClientReceive() {
	go func() {
		for {
			if ClientInstall != nil {
				buff := make([]byte,1024*1024)

				n,err := ClientInstall.Read(buff)

				msg := string(buff[:n])

				if err != nil {

					if strings.Contains(err.Error(), "use of closed network connectio") {
						fmt.Println("socket close error:", err)
						break
					} else {
						fmt.Println("receive msg error:", err)
						break
					}
					continue

				}

				fmt.Println("receive msg:", msg)

			}
		}
	}()

}
