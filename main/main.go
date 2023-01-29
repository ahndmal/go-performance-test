package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"golang.org/x/net/http2"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var httpVersion = flag.Int("version", 2, "HTTP version")

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

func GetH2(url string) {
	flag.Parse()
	client := &http.Client{}

	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	switch *httpVersion {
	case 1:
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	// Perform the request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf(
		"Got response %d: %s %s\n",
		resp.StatusCode, resp.Proto, string(body))
}

func post() {
	con, err := net.Dial("tcp", "httpbin.org:80")
	checkError(err)

	req := "POST /post HTTP/1.1\r\n" +
		"Host: httpbin.org\r\n" +
		"Content-Type: application/json\r\n\r\n" +
		"test body\r\n"

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
