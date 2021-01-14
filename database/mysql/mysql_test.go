package mysql

import (
	"github.com/panglove/BaseServer/config"
	"log"
	"testing"
)

func TestNew(t *testing.T) {

	mydb :=New(&config.MySql{
		Enable: true,
		HOST: "192.168.2.55",
		USER: "miner",
		PSWD: "Aa0011034",
		DB: "miner",
		PORT: 3306})
	log.Println(mydb.DTable("admin").DSelect([]string{"account","password"}).DWhereAnd([]string{"login_time>0"}).DExec())

	log.Println(mydb.DTable("admin").DDelete().DWhereAnd([]string{"login_time=0"}).NowSql)

	log.Println(mydb.DTable("admin").DUpdate([]string{"login_time=0"}).DWhereAnd([]string{"login_time=0"}).NowSql)

	log.Println(mydb.DTable("admin").DInsert([]string{"login_time","login_ip"},[]string{"1","'192.168.0.1'"}).DWhereAnd([]string{"login_time=0"}).DLimit([]string{"0","1"}).NowSql)
	log.Println(mydb.DTable("admin").DSelect([]string{"account","password"}).DWhereAnd([]string{"login_time>0"}).DOrderBy([]string{"id desc"}).DLimit([]string{"0","1"}).NowSql)

}