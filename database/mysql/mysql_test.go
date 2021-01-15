package mysql

import (
	"github.com/panglove/BaseServer/config"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	testString("1")
	mydb :=New(&config.MySql{
		Enable: true,
		HOST: "192.168.2.55",
		USER: "miner",
		PSWD: "Aa0011034",
		DB: "miner",
		PORT: 3306})
	log.Println(mydb.DTable("admin").DSelect("account","password").DWhereAnd("login_time>0").DExec())

	log.Println(mydb.DTable("admin").DDelete().DWhereAnd("login_time=0").NowSql)

	log.Println(mydb.DTable("admin").DUpdate("login_time=0").DWhereAnd("login_time=0").NowSql)

	log.Println(mydb.DTable("admin").DInsert([]string{"login_time","login_ip"},"1","'192.168.0.1'").DWhereAnd("login_time=0").DLimit("0","1").NowSql)
	log.Println(mydb.DTable("admin").DSelect("account","password").DWhereAnd("login_time>0").DOrderBy("id desc").DLimit("0","1").NowSql)

}

func testString(args ...string) {

	log.Println(args)

}