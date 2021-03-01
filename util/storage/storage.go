package storage

import (
	des "github.com/panglove/BaseServer/util/crypto"
	"github.com/panglove/BaseServer/util/file"
)

const StoreDir = "./data/"

func CheckDir() {

	if file.PathExists(StoreDir) {
		return
	}
	file.CreateDir(StoreDir)
}

func SetItem(key string, data string) bool {
	CheckDir()
	encodeStr, err := des.DesEncode(data)

	if err != nil {
		return false
	}

	err = file.WriteFileString(StoreDir+key+".data", encodeStr)

	return err == nil

}
func GetItem(key string) string {
	CheckDir()

	str, err := file.ReadFileString(StoreDir + key + ".data")
	if err != nil {
		return ""
	}
	str2, err := des.DesDecode(str)

	if err != nil {
		return ""
	}
	return str2
}
func RemoveItem(key string) bool {

	return file.RemovePath(StoreDir + key + ".data")

}
func RemoveAllItem() bool {

	return file.RemovePath(StoreDir)

}
