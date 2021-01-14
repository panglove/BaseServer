package database

import (
	"github.com/panglove/BaseServer/config"
	"github.com/panglove/BaseServer/database/mongodb"
	"github.com/panglove/BaseServer/database/mysql"
	"github.com/panglove/BaseServer/database/redis"
)

type DataBase struct {
	MysqlDB *mysql.MysqlDB
	MongoDB *mongodb.MongoDB
	RedisDB *redis.RedisDB
}

func New(config *config.Config) *DataBase {
	newDateBase := new(DataBase)
	newDateBase.MysqlDB = mysql.New(&config.MySql)
	newDateBase.RedisDB = redis.New(&config.Redis)
	newDateBase.MongoDB = mongodb.New(&config.MongoDB)
	return newDateBase
}
