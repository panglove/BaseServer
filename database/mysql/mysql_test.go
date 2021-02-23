package mysql

import (
	"github.com/panglove/BaseServer/config"
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
	r,R :=mydb.DTable("payment").DSelect("*").DExec()
	log.Println(r,R)

}

func testString(args ...string) {

	log.Println(args)

}