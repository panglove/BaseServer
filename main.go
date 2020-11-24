package main

import (
	"github.com/panglove/BaseServer/app"
	"github.com/panglove/BaseServer/config"
	"github.com/panglove/BaseServer/database"
	"github.com/panglove/BaseServer/net"
	"github.com/panglove/BaseServer/util/file"
	"github.com/panglove/BaseServer/util/os"
	"encoding/json"
	"log"
)

const ConfigPath = "./config.json"

var appConfig *config.Config

func main() {
	LoadConfig()
	app.Init()
	os.WaitQuit()
}
func LoadConfig() {
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
