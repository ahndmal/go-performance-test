package test

import (
	"encoding/json"
	"fmt"
	"go-http-perf/db"
	"testing"
)

func TestJson(t *testing.T) {
	for i := 0; i < 1000; i++ {
		cat := db.Cat{Name: "Murz", Age: 8, Color: "Green"}
		catJson, err := json.Marshal(cat)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(catJson))
	}
}
