package mongodb

import (
	"context"
	"fmt"
	"github.com/panglove/BaseServer/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strconv"
	"time"
)

type MongoDB struct {
	Client *mongo.Client
	Database *mongo.Database
}
func New(db *config.MongoDB)*MongoDB{
	if !db.Enable {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+db.USER+":"+db.PSWD+"@"+db.HOST+":"+ strconv.Itoa(db.PORT)+"/"+db.DB))

	if err!=nil {
		log.Println("connect mongodb error:",err)
		return nil
	}

	err = client.Ping(ctx, readpref.Primary())

	if err!=nil {
		log.Println("connect mongodb error:",err)
		return nil
	}
	newMongo :=new(MongoDB)
	newMongo.Client = client
	newMongo.Database = client.Database(""+db.DB)
	fmt.Println("Mongodb Enable Port :",db.PORT)
	return newMongo
}

func(db *MongoDB) Find(table string,filter interface{})(mr []map[string]interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection :=db.Database.Collection(table)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
		}
		mr=append(mr,result.Map())
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	return mr
}