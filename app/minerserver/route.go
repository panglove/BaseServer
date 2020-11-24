package minerserver

import (
	"github.com/panglove/BaseServer/net"
	"net/http"
)

var routeList map[string]func(http.ResponseWriter, *http.Request)

func InitRoute() {

	routeList = make(map[string]func(http.ResponseWriter, *http.Request))

	routeList["/Hello"] = OnHello

	routeList["/CheckVersion"] = OnCheckVersion


	routeList["/InsertDeviceId"] = InsertDeviceId

}
func InitWebMux() {

	for index, route := range routeList {
		net.HttpMux.HandleFunc(index, route)
	}
}
