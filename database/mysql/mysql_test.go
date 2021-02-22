package mysql

import (
	"github.com/panglove/BaseServer/config"
	"github.com/panglove/BaseServer/util/struct2"
	"log"
	"testing"
)
type Payment struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Count float64 `json:"count"`
	Time string `json:"time"`
	Share int `json:"share"`
	TotalShare int `json:"totalShare"`
	TotalEth int `json:"totalEth"`
	EthPrice int `json:"ethPrice"`
	IsPay int `json:"isPay"`
}
func TestNew(t *testing.T) {
	testString("1")
	mydb :=New(&config.MySql{
		Enable: true,
		HOST: "192.168.2.55",
		USER: "miner",
		PSWD: "Aa0011034",
		DB: "miner",
		PORT: 3306})
	newMap :=new(Payment)
	newMap.Name="hhhhh"
	log.Println(mydb.DTable("payment").DInsert(struct2.GetStructKeyList(newMap),struct2.GetStructSqlValueList(newMap)...).NowSql)
	_,R :=mydb.DTable("payment").DInsert(struct2.GetStructKeyList(newMap),struct2.GetStructSqlValueList(newMap)...).DExec()
	log.Println(R)

}

func testString(args ...string) {

	log.Println(args)

}