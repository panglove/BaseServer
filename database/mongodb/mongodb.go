package mongodb

import (
	"github.com/panglove/BaseServer/config"
	"fmt"
)

type MongoDB struct {

}
func New(db *config.MongoDB)*MongoDB{

	if !db.Enable {
		return nil
	}
	fmt.Println("Mongodb Enable Port :",db.PORT)
	return new(MongoDB)

}