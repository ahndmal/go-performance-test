package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	testFiles()
	fmt.Println(time.Now().Sub(start))
}

func testHttp() {
	limit := 30
	pass := os.Getenv("ATLAS_TOKEN")
	client := &http.Client{}
	url := fmt.Sprintf("https://marketplace.atlassian.com/rest/2/vendors/1216206/reporting/sales/transactions?limit=%d", limit)
	req, err := http.NewRequest("GET", url, nil)
	token := base64.StdEncoding.EncodeToString([]byte("quadr988@gmail.com:" + pass))
	fmt.Println(token)

	//req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Add("Authorization", "Basic "+token)

	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	bts, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bts))
}
