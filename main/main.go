package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(fmt.Sprintf("START %s", start.String()))
	fmt.Println("========")
	// ==========

	post()

	// === end
	fmt.Println(fmt.Sprintf("Script took %f seconds", time.Now().Sub(start).Seconds()))
	fmt.Println("========")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func post() {
	con, err := net.Dial("tcp", "httpbin.org:80")
	checkError(err)

	req := "POST /post HTTP/1.1\r\n" +
		"Host: httpbin.org\r\n" +
		"test body"

	_, err = con.Write([]byte(req))
	checkError(err)

	res, err := io.ReadAll(con)
	checkError(err)

	fmt.Println(string(res))
}

func head() {
	con, err := net.Dial("tcp", "webcode.me:80")
	checkError(err)

	req := "HEAD / HTTP/1.0\r\n\r\n"

	_, err = con.Write([]byte(req))
	checkError(err)

	res, err := io.ReadAll(con)
	checkError(err)

	fmt.Println(string(res))
}

func get() {
	con, err := net.Dial("tcp", "webcode.me:80")
	checkError(err)

	req := "GET / HTTP/1.1\r\n" +
		"Host: webcode.me\r\n" +
		"User-Agent: Go client\r\n\r\n"

	_, err = con.Write([]byte(req))
	checkError(err)

	res, err := ioutil.ReadAll(con)
	checkError(err)

	fmt.Println(string(res))
}
