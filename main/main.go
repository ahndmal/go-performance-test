package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"go-http-perf/docs"
	"time"
)

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

	docs.ReadDoc()
	//my_http.ConcurHttp()

	// === end
	fmt.Println(fmt.Sprintf("Script took %f seconds", time.Now().Sub(start).Seconds()))
	fmt.Println("========")
}
