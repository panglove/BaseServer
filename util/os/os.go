package os

import (
	"fmt"
	"os"
	"os/signal"
)

func WaitQuit(){

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)


	select {

	case  <-c :

		fmt.Println("Force exit BaseServer")

	}
}