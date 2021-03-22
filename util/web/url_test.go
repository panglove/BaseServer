package web

import (
	"log"
	"testing"
)

func TestParseParam(t *testing.T) {

	log.Println(ParseParam( nil))
}

func TestRequestFormPost(t *testing.T) {
	r := RequestFormPost("http://127.0.0.1/api/v1/wallet/create",Param{"hello":"test"},nil)
	log.Println(r)
}

func TestRequestJsonPost(t *testing.T) {
	r := RequestJsonPost("https://www.rchapi.com/admin/api/v1/user/login",Param{},nil)
	log.Println(r)
}