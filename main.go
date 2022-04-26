package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"go-http-perf/synch"
	"runtime"
	"sync"
	"time"
)

func main() {
	defer runtime.UnlockOSThread()
	start := time.Now()
	fmt.Println(fmt.Sprintf("START %d", start.UnixMilli()))
	fmt.Println("========")

	var wg sync.WaitGroup
	runtime.GOMAXPROCS(10)
	wg.Add(6)

	for a := 0; a < 6; a++ {
		go synch.AsyncHttp(3, &wg)
	}

	go synch.AsyncHttp(4, &wg)

	wg.Wait()

	// === end
	fmt.Println(fmt.Sprintf("Script took %d nanos", time.Now().Sub(start).Nanoseconds()))
	fmt.Println("========")
}
