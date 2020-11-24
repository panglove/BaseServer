package net

import (
	"github.com/panglove/BaseServer/util/struct2"
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseRespon struct {

}
func WebSendMsg(response http.ResponseWriter, info interface{}, Code int, Msg string) {

	if info==nil {
		info = BaseRespon{}
	}

	rebackMap := struct2.StructToEndMap(info)

	rebackMap["Code"] = Code
	rebackMap["Msg"] = Msg

	byteBuff, err := json.Marshal(rebackMap)

	if err != nil {
		fmt.Println("web send msg json err :", err)
	}
	response.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	response.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	response.Header().Set("content-type", "application/json")             //返回数据格式是json

	response.Write(byteBuff)
}
