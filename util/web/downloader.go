package web

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"github.com/panglove/BaseServer/util/file"
	"net/http"
)

func GetWebsitePath(url string, path string) string {

	return path + "/" + GetUrlHost(url)

}

func DownloadWebsiteLocal(url string, path string, webContent string) bool {

	//	webContent := GetUrlContent(url)

	if len(webContent) <= 0 {
		return false
	}
	filePath := path + "/" + GetUrlHost(url) + "/" + GetUrlFileName(url)
	file.WriteFileString(filePath, webContent)
	scriptList := GetScriptsList(webContent)

	if scriptList != nil && len(scriptList) > 0 {

		for _, link := range scriptList {
			isOk := DownloadUrl(GetAbsUrl(link, GetUrlHostHttp(url)), path)

			fmt.Println("Download ", link, " ", isOk)
		}
	}
	cssList := GetCssList(webContent)

	if cssList != nil && len(cssList) > 0 {

		for _, link := range cssList {
			isOk := DownloadUrl(GetAbsUrl(link, GetUrlHostHttp(url)), path)

			fmt.Println("Download ", link, " ", isOk)
		}
	}

	imageList := GetImageList(webContent)

	if imageList != nil && len(imageList) > 0 {

		for _, link := range imageList {
			isOk := DownloadUrl(GetAbsUrl(link, GetUrlHostHttp(url)), path)

			fmt.Println("Download ", link, " ", isOk)
		}
	}

	return true
}

func DownloadWebsite(url string, path string) bool {

	webContent := GetUrlContent(url)

	if len(webContent) <= 0 {
		return false
	}
	filePath := path + "/" + GetUrlHost(url) + "/" + GetUrlFileName(url)
	file.WriteFileString(filePath, webContent)
	scriptList := GetScriptsList(webContent)

	if scriptList != nil && len(scriptList) > 0 {

		for _, link := range scriptList {
			isOk := DownloadUrl(GetAbsUrl(link, GetUrlHostHttp(url)), path)

			fmt.Println("Download ", link, " ", isOk)
		}
	}
	cssList := GetCssList(webContent)

	if cssList != nil && len(cssList) > 0 {

		for _, link := range cssList {
			isOk := DownloadUrl(GetAbsUrl(link, GetUrlHostHttp(url)), path)

			fmt.Println("Download ", link, " ", isOk)
		}
	}

	imageList := GetImageList(webContent)

	if imageList != nil && len(imageList) > 0 {

		for _, link := range imageList {
			isOk := DownloadUrl(GetAbsUrl(link, GetUrlHostHttp(url)), path)

			fmt.Println("Download ", link, " ", isOk)
		}
	}

	return true
}

func GetUrlContent(url string) string {

	var fileBytes []byte

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		fmt.Println("request ", url, " error :", err)
		return ""
	}

	client := &http.Client{}

	resp, err := client.Do(req)

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

func DownloadUrl(url string, path string) bool {

	var fileBytes []byte

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		fmt.Println("request ", url, " error :", err)
		return false
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {

		fmt.Println("response ", url, " error :", err)
		return false
	}
	if resp.StatusCode != 200 {
		fmt.Println("response StatusCode ", url, " error :", resp.StatusCode)
		return false
	}

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(resp.Body)
		for {
			buf := make([]byte, 1024*10)
			n, err := reader.Read(buf)

			if err != nil && err != io.EOF {
				fmt.Println("read ", url, " error :", err)
				return false
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
			return false
		}

	}
	fileName := GetUrlFileName(url)
	filePath := ""
	if fileName == "index.html" {
		filePath = path + "/" + GetUrlHost(url) + "/" + GetUrlFileName(url)
	} else {

		filePath = path + "/" + GetUrlHost(url) + "/" + GetUrlPath(url)
	}

	err = file.WriteFileBytes(filePath, fileBytes)

	if err != nil {
		fmt.Println("write ", filePath, " error :", err)

		return false
	}

	return true
}
