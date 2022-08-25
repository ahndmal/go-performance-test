package html

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"testing"
)

func TestHtmlQry(t *testing.T) {
	res, err := http.Get("https://uk.wikipedia.org/wiki/%D0%9B%D0%B8%D0%BF%D0%B8%D1%86%D1%8C%D0%BA%D0%B0_%D0%B1%D0%B8%D1%82%D0%B2%D0%B0")
	if err != nil {
		log.Println("Error when getting wiki page {}", err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	doc.Find("div#mw-content-text").Each(func(i int, sel *goquery.Selection) {
		divTxt := sel.Find("div").Text()
		fmt.Println(divTxt)
	})
	doc.Find(".mw-parser-output").First().Find("a").Each(func(i int, link *goquery.Selection) {
		lText := link.Text()
		lHref, ex := link.Attr("href")
		if ex {
			fmt.Println("link is {} {}", lText, lHref)
		}
	})
	//doc.Find("h2").Each(func(i int, sel *goquery.Selection) {
	//	println(sel.Text())
	//})
}

func TestGetLinks(t *testing.T) {
	res, err := http.Get("https://www.wikipedia.org/")
	if err != nil {
		log.Println("Error when getting wiki page {}", err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	doc.Find("a").Each(func(i int, link *goquery.Selection) {
		log.Println(link.Text())
		log.Println(link.Attr("href"))
	})
}
