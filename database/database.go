package database

import (
	"github.com/panglove/BaseServer/config"
	"github.com/panglove/BaseServer/database/mongodb"
	"github.com/panglove/BaseServer/database/mysql"
	"github.com/panglove/BaseServer/database/redis"
)

func Init(config *config.Config){
	redis.Init(&config.Redis)
	mongodb.Init(&config.MongoDB)
	mysql.Init(&config.MySql)
}