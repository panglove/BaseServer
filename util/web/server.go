package web

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

var CurrSvr *http.Server

func HttpFileServer(port int, path string) {
	if CurrSvr != nil {
		CurrSvr.Shutdown( context.Background())
	}

	newMyx := http.NewServeMux()

	newMyx.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))

	CurrSvr = &http.Server{Addr: ":" + strconv.Itoa(port), Handler: newMyx}

	go func() {

		err := CurrSvr.ListenAndServe()

		if err != nil {

			fmt.Println("http listen error :", err)
		}
	}()

}
