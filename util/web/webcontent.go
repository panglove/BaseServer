package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetScriptsList(content string) map[int]string {

	linkMap := make(map[int]string)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {

		fmt.Println("async html content error:", err)
		return nil

	}
	count := 0
	doc.Find("script").Each(func(i int, selection *goquery.Selection) {
		link, isExist := selection.Attr("src")
		if isExist {
			linkMap[count] = link
			count++
		}
	})

	if count > 0 {
		return linkMap
	} else {
		return nil
	}

}
func GetImageList(content string) map[int]string {

	linkMap := make(map[int]string)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {

		fmt.Println("async html content error:", err)
		return nil

	}
	count := 0
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		link, isExist := selection.Attr("src")
		if isExist {
			linkMap[count] = link
			count++
		}
	})

	if count > 0 {
		return linkMap
	} else {
		return nil
	}

}
func GetCssList(content string) map[int]string {

	linkMap := make(map[int]string)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {

		fmt.Println("async html content error:", err)
		return nil

	}
	count := 0
	doc.Find("link").Each(func(i int, selection *goquery.Selection) {
		linkType, isExist := selection.Attr("rel")
		if isExist {


			if linkType == "stylesheet" {

				linkRel, isExist := selection.Attr("href")

				if isExist {

					linkMap[count] = linkRel
					count++
				}

			}
		}
	})

	if count > 0 {
		return linkMap
	} else {
		return nil
	}
}
