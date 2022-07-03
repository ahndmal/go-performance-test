package html

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"os"
	"strings"
	"testing"
)

func TestQueryOne(t *testing.T) {
	//s := `<html><body></body></html>`
	htmlFIle, err := os.ReadFile("test1.html")
	doc, err := htmlquery.Parse(strings.NewReader(string(htmlFIle)))
	fmt.Println(doc)
	if err != nil {
		fmt.Println(err)
	}
	divs := htmlquery.Find(doc, "//div")
	fmt.Println(divs)
}

func example() {
	doc, err := htmlquery.LoadURL("https://www.bing.com/search?q=golang")
	if err != nil {
		panic(err)
	}
	// Find all news item.
	list, err := htmlquery.QueryAll(doc, "//ol/li")
	if err != nil {
		panic(err)
	}
	for i, n := range list {
		a := htmlquery.FindOne(n, "//a")
		if a != nil {
			fmt.Printf("%d %s(%s)\n", i, htmlquery.InnerText(a), htmlquery.SelectAttr(a, "href"))
		}
	}
}
