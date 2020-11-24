package sql

import (
	"github.com/panglove/BaseServer/database/mysql"
)
func IsDeviceIdExist(deviceId string) bool {

	deviceList,isOk := mysql.Select("select * from deviceList where deviceId = ?",deviceId)

	if isOk && len(deviceList)>0{
		return true
	}
	return false


}
func InsertDeviceId(deviceId string,sunId string) bool {
	_,err :=mysql.MysqlDB.Exec("insert into deviceList (deviceId,sunId) values (?,?)",deviceId,sunId)

	return err==nil
}

func IsSunIdExist(sunId string) bool {

	deviceList,isOk := mysql.Select("select * from deviceList where sunId = ?",sunId)

	if isOk && len(deviceList)>0{
		return true
	}
	return false


}