package my_http

import (
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
)

func TestHttpLocal(t *testing.T) {
	for a := 0; a < 10000; a++ {
		response, err := http.Get("http://localhost:2011")
		if err != nil {
			log.Println(err)
		}
		readAll, err2 := io.ReadAll(response.Body)
		if err2 != nil {
			log.Println(err2)
		}
		log.Println(string(readAll))
	}
}

func TestHttpLocalAsync(t *testing.T) {
	wg := sync.WaitGroup{}
	for a := 0; a < 10000; a++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			response, err := http.Get("http://localhost:2011")
			if err != nil {
				log.Println(err)
			}
			readAll, err2 := io.ReadAll(response.Body)
			if err2 != nil {
				log.Println(err2)
			}
			log.Println(string(readAll))
		}()
	}
	wg.Wait()

}
