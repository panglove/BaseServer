package net

import "github.com/panglove/BaseServer/config"

func Init(config *config.Config) {
	InitWeb(&config.WebServer)
	InitSocket(&config.SocketServer)
	InitWebSocket(&config.WebSocketServer)

}