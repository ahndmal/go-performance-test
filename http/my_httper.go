package my_http

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func PerformReq(url string, times int, wg *sync.WaitGroup) {
	//defer wg.Done()
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error when performing GET HTTP. %s", err)
	}
	//for i := 0; i < times; i++ {
	response, err2 := client.Do(req)
	if err2 != nil {
		log.Panic(err)
	}
	body, err3 := io.ReadAll(response.Body)
	if err3 != nil {
		return
	}
	fmt.Println(string(body))
	//}

}

func ConcurHttp() {
	defer runtime.UnlockOSThread()

	var wg sync.WaitGroup
	runtime.GOMAXPROCS(100)

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
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(body))

		}()
	}

	wg.Wait()
}
