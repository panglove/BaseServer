package sqlserver

import (
	"github.com/panglove/BaseServer/config"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	testString("1")
	mydb := New(&config.MsSql{
		Enable: true,
		HOST:   "47.57.15.159",
		USER:   "sa",
		PSWD:   "JSs7GihgoKM8rZQm",
		DB:     "RCH",
		PORT:   1433})
	maps, isok := mydb.DTable("Machine").DSelect("*").DWhereAnd("Address is not null ", " Address <> ''").DExec()

	if isok {
		for k, _ := range maps {

			log.Println(maps[k])
		}
	}

	//log.Println(mydb.DTable("admin").DDelete().DWhereAnd("login_time=0").NowSql)
	//
	//log.Println(mydb.DTable("admin").DUpdate("login_time=0").DWhereAnd("login_time=0").NowSql)
	//
	//log.Println(mydb.DTable("admin").DInsert([]string{"login_time", "login_ip"}, "1", "'192.168.0.1'").DWhereAnd("login_time=0").DLimit("0", "1").NowSql)
	//log.Println(mydb.DTable("admin").DSelect("account", "password").DWhereAnd("login_time>0").DOrderBy("id desc").DLimit("0", "1").NowSql)

}

func testString(args ...string) {

	log.Println(args)

}
