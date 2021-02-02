package config

var AppConifg *Config

type MongoDB struct {
	Enable bool
	HOST   string
	USER   string
	PSWD   string
	DB     string
	PORT   int
}
type MySql struct {
	Enable bool
	HOST   string
	USER   string
	PSWD   string
	DB     string
	PORT   int
}
type MsSql struct {
	Enable bool
	HOST   string
	USER   string
	PSWD   string
	DB     string
	PORT   int
}

type Redis struct {
	Enable bool
	HOST   string
	USER   string
	PSWD   string
	DB     string
	PORT   int
}

type WebServer struct {
	Enable     bool
	PORT       int
	UseStatic  bool
	StaticPath string
	Host string
}
type WebSocketServer struct {
	Enable bool
	PORT   int
}

type SocketServer struct {
	Enable bool
	PORT   int
}

type Config struct {
	MongoDB         MongoDB
	MySql           MySql
	MsSql           MsSql
	Redis           Redis
	WebServer       WebServer
	WebSocketServer WebSocketServer
	SocketServer    SocketServer
}
