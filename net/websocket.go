package net

import (
	"github.com/panglove/BaseServer/config"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strconv"
	"strings"
)
var WebsocketMux *http.ServeMux

type WebsocketUtil interface {
	OnMessage(wsCon *websocket.Conn, msg string)
	OnError(wsCon *websocket.Conn, msg string)
}

var WebsocketReciverList map[int]WebsocketUtil

func AddWebsocketReciver(reciver WebsocketUtil) {

	WebsocketReciverList[len(WebsocketReciverList)] = reciver

}
func InitWebSocket(cf *config.WebSocketServer) {
	if !cf.Enable {
		return
	}
	fmt.Println("WebSocket Enable Port :", cf.PORT, "/ws")

	WebsocketReciverList = make(map[int]WebsocketUtil)
	WebsocketMux := http.NewServeMux()
	WebsocketMux.Handle("/ws", websocket.Handler(OnWebsocketHander))
	WebsocketMux.HandleFunc("/", OnWebsocketIndexHander)

	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(cf.PORT), WebsocketMux)

		if err != nil {
			log.Fatal("listen websocket error:", err)
		}
	}()
}
func OnWebsocketHander(wsCon *websocket.Conn) {
	GoWebsocketReceiveMsg(wsCon)
}
func OnWebsocketIndexHander(resp http.ResponseWriter, rq *http.Request) {
	resp.WriteHeader(200)
	resp.Write([]byte("welcome game websocket server!"))
}
func GoWebsocketReceiveMsg(wsCon *websocket.Conn) {

	for {
		var msg string

		err := websocket.Message.Receive(wsCon, &msg)

		if err == nil {
		//	fmt.Println("connect msg :" + msg)

			for _, receiver := range WebsocketReciverList {
				receiver.OnMessage(wsCon, msg)
			}
		} else {
			for _, receiver := range WebsocketReciverList {
				receiver.OnError(wsCon, err.Error())
			}
			if strings.Contains(err.Error(), "use of closed network connectio") {
				fmt.Println("socket close error:", err)
			} else {
				fmt.Println("receive msg error:", err)
			}
			return

		}
	}

}

func SendWebsocketMsg(wsCon *websocket.Conn, msg string) {
	go func() {
		if wsCon != nil {
			err := websocket.Message.Send(wsCon, msg)

			if err != nil {
				fmt.Println("send msg error:", err)
			}
		}
	}()
}

func SendWebsocketInterface(wsCon *websocket.Conn, msg interface{}) {
	go func() {
		if wsCon != nil {
			err := websocket.Message.Send(wsCon, msg)

			if err != nil {
				fmt.Println("send msg error:", err)
			}
		}
	}()
}
func SendWebsocketProtocoInterface(wsCon *websocket.Conn, head string, info string) {
	go func() {
		if wsCon != nil {

			sendStr, err2 := json.Marshal(Protoco{head, info})
			if err2 != nil {
				fmt.Println("send msg error:", err2)
			}
			err := websocket.Message.Send(wsCon, string(sendStr))

			if err != nil {
				fmt.Println("send msg error:", err)
			} else {

				fmt.Println("send msg :", string(sendStr))

			}
		}
	}()
}
