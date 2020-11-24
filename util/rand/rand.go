package rand

import (
	"crypto/rand"
	"math/big"
	"reflect"
)

func GetRandRange(min int64, max int64) int64 {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(max))
	if randInt.Cmp(big.NewInt(min)) < 0 {
		randInt = big.NewInt(GetRandRange(min, max))
	}
	return randInt.Int64()
}
func GetRandArray(randCount int, array interface{}) []interface{} {

	var randIntList []int64

	var newList []interface{}


	newArray := SliceToInterface(array)

	for k := 0; k < randCount; k++ {
		randInt := int64(0)

		for {
			randInt = GetRandRange(0, int64(len(newArray)))
			isHas := false

			for _, i := range randIntList {
				if randInt == i {
					isHas = true
				}
			}
			if !isHas {
				break
			}

		}
		randIntList = append(randIntList, randInt)
	}

	for o := 0; o < len(randIntList); o++ {
		newList = append(newList, newArray[randIntList[o]])

	}
	return newList

}
func SliceToInterface(slice interface{}) []interface{} {

	value := reflect.ValueOf(slice)

	len := value.Len()

	list := make([]interface{}, len)

	for i := 0; i < len; i++ {
		list[i] = value.Index(i).Interface()
	}
	return list
}
