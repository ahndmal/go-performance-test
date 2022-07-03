package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"go-http-perf/dirs"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(fmt.Sprintf("START %s", start.String()))
	fmt.Println("========")

	println(dirs.CountWords(dirs.GetWholeText()))
	fileBytes, err := ioutil.ReadFile("/home.andrii/docs/file1.txt")
	if err != nil {
		println(err)
	}
	count := 0
	for _, w := range strings.Split(string(fileBytes), " ") {
		println(w)
		count += 1
	}
	fmt.Println(count)

	// === end
	fmt.Println(fmt.Sprintf("Script took %f seconds", time.Now().Sub(start).Seconds()))
	fmt.Println("========")
}

func ConcurHttp() {
	defer runtime.UnlockOSThread()

	var wg sync.WaitGroup
	runtime.GOMAXPROCS(10)

	links := []string{
		"https://en.wikipedia.org/wiki/1921%E2%80%9322_Cardiff_City_F.C._season",
		"https://en.wikipedia.org/wiki/Football_League_Second_Division",
		"https://en.wikipedia.org/wiki/William_McGregor_(football)",
		"https://en.wikipedia.org/wiki/Perthshire",
		"https://en.wikipedia.org/wiki/Clackmannanshire",
		"https://en.wikipedia.org/wiki/Scottish_Gaelic",
		"https://en.wikipedia.org/wiki/2022_Slovenian_parliamentary_election",
	}
	wg.Add(len(links))

	for _, l := range links {
		link := l
		go func() {
			defer wg.Done()
			response, err := http.Get(link)
			if err != nil {
				fmt.Println(err)
			}
			//body, err := ioutil.ReadAll(response.Body)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//println(string(body))
			doc, err := html.Parse(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			println(doc.Parent)

		}()
	}

	wg.Wait()
}
