package test

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func testHttp() {
	limit := 30
	pass := os.Getenv("ATLAS_TOKEN")
	client := &http.Client{}
	url := fmt.Sprintf("%s/transactions?limit=%d", os.Getenv("MKT_LINK"), limit)
	req, err := http.NewRequest("GET", url, nil)
	token := base64.StdEncoding.EncodeToString([]byte(os.Getenv("USER_MAIL") + ":" + pass))

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
