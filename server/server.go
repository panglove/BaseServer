package server

import (
	"github.com/panglove/BaseServer/config"
	"github.com/panglove/BaseServer/util/os"
)

var appConfig *config.Config

type BaseApp interface {
	Init()
}
func Init(ConfigPath string,app BaseApp) {
	app.Init()
	os.WaitQuit()
}
