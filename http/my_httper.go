package my_http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func PerformReq(url string, times int, wg *sync.WaitGroup) {
	wg.Add(1)
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error when performing GET HTTP. %s", err)
	}
	response, err2 := client.Do(req)
	if err2 != nil {
		log.Panic(err)
	}
	body, err3 := io.ReadAll(response.Body)
	if err3 != nil {
		return
	}

	fmt.Println(string(body))
	wg.Done()
}
