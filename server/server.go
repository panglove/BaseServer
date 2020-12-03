package server

import (
	"encoding/json"
	"github.com/panglove/BaseServer/config"
	"github.com/panglove/BaseServer/database"
	"github.com/panglove/BaseServer/net"
	"github.com/panglove/BaseServer/util/file"
	"github.com/panglove/BaseServer/util/os"
	"log"
)

var appConfig *config.Config

type BaseApp interface {
	Init()
}
func Init(ConfigPath string,app BaseApp) {
	LoadConfig(ConfigPath)
	app.Init()
	os.WaitQuit()
}
func LoadConfig(ConfigPath string) {
	configStr, err := file.ReadFileString(ConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	appConfig = new(config.Config)

	json.Unmarshal([]byte(configStr), appConfig)

	if err != nil {
		log.Fatal(err)
	}
	config.AppConifg = appConfig
	database.Init(appConfig)
	net.Init(appConfig)

}
