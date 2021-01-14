package redis

import "github.com/panglove/BaseServer/config"

type RedisDB struct {

}
func New(db *config.Redis)*RedisDB{

	if !db.Enable {
		return nil
	}
	return new(RedisDB)
}
