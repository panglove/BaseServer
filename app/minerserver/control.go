package minerserver

import (
	"github.com/panglove/BaseServer/app/minerserver/protoco"
	"github.com/panglove/BaseServer/app/minerserver/sql"
	"github.com/panglove/BaseServer/net"
	"github.com/panglove/BaseServer/util/file"
	"encoding/json"
	"fmt"
	"net/http"
)

func OnHello(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("this is a all base server!"))

}
func OnCheckVersion(response http.ResponseWriter, request *http.Request) {
	versionPath := "./static/version.json"

	versionContent, err := file.ReadFileString(versionPath)

	versionInfo := new(protoco.VersionInfo)

	err2 := json.Unmarshal([]byte(versionContent), versionInfo)

	if err != nil || err2 != nil {
		fmt.Println("version file not exists")
		response.WriteHeader(400)
		net.WebSendMsg(response, protoco.VersionInfo{0, ""}, -1, "")
		return
	}
	net.WebSendMsg(response, *versionInfo, 1, "ok")

}
func GetNewStartSh(response http.ResponseWriter, request *http.Request) {
	versionPath := "./static/version.json"

	versionContent, err := file.ReadFileString(versionPath)

	versionInfo := new(protoco.VersionInfo)

	err2 := json.Unmarshal([]byte(versionContent), versionInfo)

	if err != nil || err2 != nil {
		fmt.Println("version file not exists")
		response.WriteHeader(400)
		net.WebSendMsg(response, protoco.VersionInfo{0, ""}, -1, "")
		return
	}
	net.WebSendMsg(response, *versionInfo, 1, "ok")

}

func InsertDeviceId(response http.ResponseWriter, request *http.Request) {

	keys := request.URL.Query()

	 deviceId := keys.Get("deviceId")

	 sunId := keys.Get("sunId")

	 if len(sunId) ==0 || len(deviceId) == 0 {

		 net.WebSendMsg(response, nil, -1, "设备ID和向日葵ID不能为空")
		 return

	 }
	 if sql.IsDeviceIdExist(deviceId) {
		 net.WebSendMsg(response, nil, -1, "设备ID已存在")
		 return

	 }
	if sql.IsSunIdExist(sunId) {
		net.WebSendMsg(response, nil, -1, "向日葵ID已存在")
		return
	}

	 fmt.Println("插入设备:",deviceId,",",sunId)

	 isOk :=sql.InsertDeviceId(deviceId,sunId)

	 if isOk {

		 net.WebSendMsg(response, nil, 1, "插入成功")

	 }else {

		 net.WebSendMsg(response, nil, -1, "插入失败")

	 }


}
