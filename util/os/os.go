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

	case s := <-c :
		fmt.Println("Got signal:", s)

	}
}