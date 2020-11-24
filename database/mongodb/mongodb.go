package mongodb

import (
	"github.com/panglove/BaseServer/config"
	"fmt"
)

func Init(db *config.MongoDB){

	if !db.Enable {
		return
	}
	fmt.Println("Mongodb Enable Port :",db.PORT)

}