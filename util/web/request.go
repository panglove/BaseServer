package web

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


// http://www.baidu.com {"key":1,"key2":2}
func RequestGet(url string,param map[string]interface{}) string {

	var fileBytes []byte

	client := &http.Client{Timeout:time.Second*15}

	reqParam :=ParseParam(param)

	if len(reqParam)>0 {
		url = url+"?"+reqParam
	}
	resp, err := client.Get(url)

	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {

		fmt.Println("response ", url, " error :", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("response StatusCode ", url, " error :", resp.StatusCode)
		return ""
	}

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(resp.Body)
		for {
			buf := make([]byte, 1024*10)
			n, err := reader.Read(buf)

			if err != nil && err != io.EOF {
				fmt.Println("read ", url, " error :", err)
				return ""
			}

			if n == 0 {
				break
			}
			fileBytes = append(fileBytes, buf...)
		}
	default:
		fileBytes, err = ioutil.ReadAll(resp.Body)

		if err != nil && err != io.EOF {
			fmt.Println("read ", url, " error :", err)
			return ""
		}

	}
	return string(fileBytes)
}
// http://www.baidu.com {"key":1,"key2":2}
func RequestJsonPost(url string,param map[string]interface{},headers map[string]string) string {
	reqParam ,err :=json.Marshal(param)

	if err != nil {
		fmt.Println("parame parse ", param, " error :", err)
		return ""
	}
	newReq,err:=http.NewRequest("POST",url,strings.NewReader(string(reqParam)))
	if err != nil {
		fmt.Println("req ", url, " error :", err)
		return ""
	}
	newReq.Header.Set("Content-Type","application/json")
	for h:=range headers {
		newReq.Header.Set(h,headers[h])
	}
	var fileBytes []byte

	client := &http.Client{Timeout:time.Second*15}

	resp, err := client.Do(newReq)

	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {

		fmt.Println("response ", url, " error :", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("response StatusCode ", url, " error :", resp.StatusCode)
		return ""
	}

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(resp.Body)
		for {
			buf := make([]byte, 1024*10)
			n, err := reader.Read(buf)

			if err != nil && err != io.EOF {
				fmt.Println("read ", url, " error :", err)
				return ""
			}

			if n == 0 {
				break
			}
			fileBytes = append(fileBytes, buf...)
		}
	default:
		fileBytes, err = ioutil.ReadAll(resp.Body)

		if err != nil && err != io.EOF {
			fmt.Println("read ", url, " error :", err)
			return ""
		}

	}
	return string(fileBytes)
}

// http://www.baidu.com {"key":1,"key2":2}
func RequestFormPost(url string,param map[string]interface{},headers map[string]string) string {
	reqParam :=ParseParam(param)
	newReq,err:=http.NewRequest("POST",url,strings.NewReader(reqParam))
	if err != nil {
		fmt.Println("req ", url, " error :", err)
		return ""
	}
	newReq.Header.Set("Content-Type","application/x-www-form-urlencoded")
	for h:=range headers {
		newReq.Header.Set(h,headers[h])
	}
	var fileBytes []byte

	client := &http.Client{Timeout:time.Second*15}

	resp, err := client.Do(newReq)

	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {

		fmt.Println("response ", url, " error :", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("response StatusCode ", url, " error :", resp.StatusCode)
		return ""
	}

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(resp.Body)
		for {
			buf := make([]byte, 1024*10)
			n, err := reader.Read(buf)

			if err != nil && err != io.EOF {
				fmt.Println("read ", url, " error :", err)
				return ""
			}

			if n == 0 {
				break
			}
			fileBytes = append(fileBytes, buf...)
		}
	default:
		fileBytes, err = ioutil.ReadAll(resp.Body)

		if err != nil && err != io.EOF {
			fmt.Println("read ", url, " error :", err)
			return ""
		}

	}
	return string(fileBytes)
}