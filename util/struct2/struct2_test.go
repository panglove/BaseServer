package struct2

import (
	"encoding/json"
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
func TestGetStructKeyList(t *testing.T) {
	newStruct :=new(Payment)
	newStruct.Time ="helloworld"
	log.Println(GetStructKeyList(newStruct))
}

func TestMapToToEndStruct(t *testing.T) {
	newMap:=make(map[string]interface{})
	newMap["id"] = "1"
	newMap["Time"] = "hello world"
	newMap["count"] = "19.2265"
	newStruct :=new(Payment)
	buff :=GetMapToStructBuff(newMap,*newStruct)
	json.Unmarshal(buff,newStruct)
	log.Println(*newStruct)


}