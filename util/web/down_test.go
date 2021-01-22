package web

import (
	"fmt"
	"github.com/panglove/BaseServer/util/file"
	"testing"
)

func TestHttpFileServer(t *testing.T) {
	HttpFileServer(8082, "./static/www.tronsix.com")
	select {}
}

func TestDown(t *testing.T) {
	precacheConfig := [][]string{{"./index.html", "ff77117fa490b06c8591327b3cdb74e9"}, {"./static/css/main.6b255797.css", "6b255797786234ccc6fcededcc560660"}, {"./static/js/0.fd57a453.chunk.js", "afb95edbb3189ff9d25136289fcda4ab"}, {"./static/js/1.7262e8ac.chunk.js", "62954c5a37da878c2fd7e025f829f9f0"}, {"./static/js/main.73dfaa5b.js", "eb6c7d681761e1a5ea929c15ee5698e0"}, {"./static/media/1.a94e77b4.png", "a94e77b4531e07801e1465ee2dd2191c"}, {"./static/media/2.5082f392.png", "5082f392e17ce2318e6e10b3f66211c7"}, {"./static/media/3.05e449e2.png", "05e449e2f71fcf91a866a61ec22a688f"}, {"./static/media/4.ab27c600.png", "ab27c600f842dfd85c229f8c48f8f7bf"}, {"./static/media/404.651005df.jpg", "651005df135c79d2679c8a1bf39c5066"}, {"./static/media/5.85a59683.png", "85a59683d6e7fee94ab34e8890364291"}, {"./static/media/6.0a313b0e.png", "0a313b0ee0b48fef5d3f9a8f287875e7"}, {"./static/media/arrow.3ac027f1.png", "3ac027f17f2cfbc1b3d0024fb30213ef"}, {"./static/media/bet1.7f0a5eb2.png", "7f0a5eb26f5789066ae931c611ab2f96"}, {"./static/media/bet2.f5e08a16.png", "f5e08a16d0bd47bba5e402b53f9d9444"}, {"./static/media/bet3.c6ce0fc6.png", "c6ce0fc670f390a3f811ede16f68b0ad"}, {"./static/media/bet4.6c3c2c13.png", "6c3c2c1322dee58efcf740085880f805"}, {"./static/media/bet5.3f13c57f.png", "3f13c57f72fa588ef08dff9bab1a43cc"}, {"./static/media/bet6.6b75c88d.png", "6b75c88de54bb23bf32ab6f0edddd51c"}, {"./static/media/bg.e65228bb.png", "e65228bb054b37cb15cdc48c7900eba3"}, {"./static/media/bi.148b047e.png", "148b047ee4f35fdf4ff1e4351103dda0"}, {"./static/media/close.74705907.png", "747059070424976bbf175484e9b23104"}, {"./static/media/dajiang.42b07304.png", "42b0730494525d5a719f31a1af310bf7"}, {"./static/media/dianbao.f8454040.png", "f84540404787a7ad921576b092b347b5"}, {"./static/media/dice-loading.ef8e2b70.gif", "ef8e2b70ec018ee3877b5cbe5ef316bc"}, {"./static/media/left-logo.e7784e6a.png", "e7784e6aee3ed49d02e1400c7a619c27"}, {"./static/media/loading.68544878.png", "68544878a363cf653dc3f0f0d13a39d2"}, {"./static/media/logo.9f929336.png", "9f92933603a10e2d862d1a33dd8a6962"}, {"./static/media/mingzi.0952b70c.png", "0952b70ce6c3693a45925e989870713d"}, {"./static/media/page-bg.c88fa709.jpg", "c88fa7090366e653f53015327996046c"}, {"./static/media/popup_bottom.7da15832.png", "7da1583263e2ba9081a2ff5493cb5a50"}, {"./static/media/popup_top.8539ee2d.png", "8539ee2d85cc3175f831483012eadb11"}, {"./static/media/qukuai.5d041766.png", "5d04176607a2cd2512972ae2c7476092"}, {"./static/media/result1.ef42f345.png", "ef42f345ba5232e54668f64b54fe8d21"}, {"./static/media/result2.bbdfd644.png", "bbdfd644a9593501afe7a7231955a407"}, {"./static/media/result3.5fbf012b.png", "5fbf012bc42cb4726178809d829ea9b2"}, {"./static/media/result4.77602a88.png", "77602a88223f3b08e40bd7f592f31f21"}, {"./static/media/result5.f6a00be5.png", "f6a00be54026b791ee77ea64f130220d"}, {"./static/media/result6.745600ca.png", "745600cadbf7eaa95b3949e16d62d3cf"}, {"./static/media/six_bg.f1d2e701.png", "f1d2e7010f0bb7517eca82d0e9546978"}, {"./static/media/tuite.2ee9a01f.png", "2ee9a01ffab9593c5f70fc8dc3f8e387"}, {"./static/media/w1.ab95bd39.png", "ab95bd39ac8383e80c4b4686b30069cb"}, {"./static/media/w2.7ae6cb56.png", "7ae6cb5681d7ead446870e0e8b87a808"}, {"./static/media/w3.6a17cde7.png", "6a17cde7da12e6617fbe845e78be7f33"}, {"./static/media/w4.db11c52b.png", "db11c52b068fce23d7461d2be73f7453"}, {"./static/media/w5.0a02f8a4.png", "0a02f8a4a4f597003ed8784c9cd23c0f"}, {"./static/media/w6.54760512.png", "5476051260214eb65d1d15ba3867613e"}}

	for i := 0; i < len(precacheConfig); i++ {

		isok := DownloadUrl("http://www.tronsix.com/"+precacheConfig[i][0], "./static")
		fmt.Println(isok)
	}
}
func TestUrl(t *testing.T) {
	fmt.Println("get ", GetUrlFileName("http://www.ethereumx.xyz/MyCrypto-ForETX-Setup1.7.9.exe"))
}
func TestUrlHost(t *testing.T) {
	fmt.Println("get ", GetUrlHost("https://www.ethereumx.xyz/"))
}

func TestGetUrlContent(t *testing.T) {
	content := GetUrlContent("http://www.tronsix.com")
	fmt.Println(content)

	if len(content) > 0 {

		linkMap := GetCssList(content)

		fmt.Println(linkMap)

	}
}
func TestGetUrlScriptContent(t *testing.T) {
	content := GetUrlContent("http://www.tronsix.com")
	fmt.Println(content)

	if len(content) > 0 {

		linkMap := GetScriptsList(content)

		fmt.Println(linkMap)

	}
}

func TestDownloadWebsite(t *testing.T) {
	localHtml, _ := file.ReadFileString("./index.html")
	isok := DownloadWebsiteLocal("http://www.tronsix.com", "./static", localHtml)
	fmt.Println(isok)
	HttpFileServer(8082, "./static/www.tronsix.com")
	select {}
}

func TestWebsite(t *testing.T) {
	HttpFileServer(8082, "./static/www.tronsix.com")
	select {}
}
func TestDownloadWebsite2(t *testing.T) {
	isok := DownloadWebsite("https://simple-yfi.finance/pool1.html", "./static")
	fmt.Println(isok)
	HttpFileServer(8082, "./static/simple-yfi.finance")
	select {}
}