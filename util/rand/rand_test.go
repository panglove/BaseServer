package rand

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {


	t.Log(GetRandRange(10,100))

}

func TestB(t *testing.T) {

	randArray :=  []string{"1","dojksapokd","djdjksj","djdjh","djdj"}


	rangeArray := GetRandArray(5,randArray)

	fmt.Println(rangeArray)





}
