package net

import (
	"github.com/panglove/BaseServer/config"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestWeb(ct *testing.T) {

	cf := config.WebServer{PORT: 8081, UseStatic: false, StaticPath: "static", Enable: true}

	if !cf.Enable {
		return
	}


	if cf.UseStatic {

		fsh := http.FileServer(http.Dir(cf.StaticPath))
		http.Handle("/", http.StripPrefix("/", fsh))

		fmt.Println("web use static path ", cf.StaticPath)
	} else {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

			writer.WriteHeader(200)
			writer.Write([]byte("this web root path"))

		})

	}
	go func() {
		fmt.Println("Web Enable Port :", cf.PORT)
		mux := http.NewServeMux()


		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

			writer.WriteHeader(200)
			writer.Write([]byte("this web root path 2"))

		})
		err := http.ListenAndServe(":"+strconv.Itoa(cf.PORT), mux)

		if err != nil {
			log.Fatal("listen web error:", err)
		}

	}()
	fmt.Println("Web Enable Port :", cf.PORT+1)
	err := http.ListenAndServe(":"+strconv.Itoa(cf.PORT+1), nil)
	if err != nil {
		log.Fatal("listen web error:", err)
	}
}
