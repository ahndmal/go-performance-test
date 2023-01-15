package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	my_http "go-http-perf/http"
	"io/ioutil"
	"strings"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	start := time.Now()
	fmt.Println(fmt.Sprintf("START %s", start.String()))
	fmt.Println("========")
	// ======

	//wg := sync.WaitGroup{}
	//url := "https://en.wikipedia.org/wiki/Shannon_Lucid"
	//for i := 0; i < 50; i++ {
	//	go func() {
	//		wg.Add(1)
	//		my_http.PerformReq(url, 50, &wg)
	//		wg.Done()
	//	}()
	//}

	//docs.ReadDoc2()
	my_http.ConcurHttp()

	// === end
	fmt.Println(fmt.Sprintf("Script took %f seconds", time.Now().Sub(start).Seconds()))
	fmt.Println("========")
}

func countWords2() {
	fileBytes, err := ioutil.ReadFile("/home/andrii/Documents/lorem1.txt")
	if err != nil {
		println(err)
	}
	count := 0
	for _, w := range strings.Split(string(fileBytes), " ") {
		println(w)
		count += 1
	}
	fmt.Println(count)
}
