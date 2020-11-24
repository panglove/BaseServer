package redis

import "github.com/panglove/BaseServer/config"

func Init(db *config.Redis){

	if !db.Enable {
		return
	}


}
