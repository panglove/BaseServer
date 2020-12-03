golang实现的基础服务器框架，简单支持http，websocket 协议等。支持连接多种数据库。封装常用加密方法，文件处理方法，随机数，数据储存方法等。
详细使用例子参考./test/test.go  
或者克隆整个项目去运行


简单HTTP搭建

package test  

import (    
	"fmt"  
	"github.com/panglove/BaseServer/net"  
	"github.com/panglove/BaseServer/server"  
	"net/http"  
)

type BaseServer struct {

}  
var baseWebApp BaseServer  

func Start()  {    
	server.Init("./config.json",baseWebApp)    
}    
    
//app install implement    
func(server BaseServer) Init(){    
	fmt.Println("launch http server")  
	TestWebRoute()  
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
  

其中config.json文件  
  
{  
  "MongoDB": {  
    "Enable": false,  
    "HOST": "127.0.0.1",  
    "USER": "db",  
    "PSWD": "dbpass",  
    "DB": "dbDB",  
    "PORT": 27017  
  },  
  "MySql": {  
    "Enable": true,  
    "HOST": "127.0.0.1",  
    "USER": "db",  
    "PSWD": "dbpass",  
    "DB": "db",  
    "PORT": 3306  
  },    
  "Redis": {  
    "Enable": false,  
    "HOST": "127.0.0.1",  
    "USER": "db",  
    "PSWD": "dbpass",  
    "DB": "dbDB",  
    "PORT": 27017  
  },  
  "WebServer": {  
    "Enable": true,  
    "UseStatic": true,  
    "StaticPath": "static",  
    "Host": "http://0.0.0.0:8567",  
    "PORT": 8567  
  },  
  "WebSocketServer": {  
    "Enable": true,  
    "PORT": 8083  
  },  
  "SocketServer": {  
    "Enable": false,  
    "PORT": 8083  
  }  
}  
  
设置好后，自动开启相应服务和开启数据库。