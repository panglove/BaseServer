package net

import (
	"github.com/panglove/BaseServer/config"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var HttpMux *http.ServeMux

func InitWeb(cf *config.WebServer) {
	if !cf.Enable {
		return
	}
	fmt.Println("Web Enable Port :", cf.PORT)

	HttpMux = http.NewServeMux()

	if cf.UseStatic {

		fsh := http.FileServer(http.Dir(cf.StaticPath))

		HttpMux.Handle("/static/", StripPrefix("/static/", fsh))
		fmt.Println("web use static path ", cf.StaticPath)
	}
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(cf.PORT), HttpMux)
		if err != nil {
			log.Fatal("listen web error:", err)
		}
	}()

}
func StripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}
