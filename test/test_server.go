package test

import (
	"fmt"
	"github.com/panglove/BaseServer/net"
	"github.com/panglove/BaseServer/server"
	"golang.org/x/net/websocket"
	"net/http"
)
type WebsocketServer struct {

}
type BaseServer struct {

}
var baseWebApp BaseServer
var baseSocketApp WebsocketServer

func Start()  {
	server.Init("./config.json",baseWebApp)
}

//app install implement
func(server BaseServer) Init(){

	fmt.Println("launch http server")

	TestWebRoute()
	TestWebsocket()
}
//web server test
func TestWebRoute()  {
	net.HttpMux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		net.WebSendMsg(writer,nil,1,"true")
	})

	net.HttpMux.HandleFunc("/test2", func(writer http.ResponseWriter, request *http.Request) {
		net.WebSendMsg(writer,nil,-1,"false")
	})

	
}

func TestWebsocket()  {
	net.AddWebsocketReciver(baseSocketApp)
}

//websocket msg reciver
func(ws WebsocketServer) OnMessage(wsCon *websocket.Conn, msg string)  {

	fmt.Println("receive:",msg)

	net.SendWebsocketMsg(wsCon,"hello client")

}
//websocket error
func(ws WebsocketServer) OnError(wsCon *websocket.Conn, msg string)  {

}