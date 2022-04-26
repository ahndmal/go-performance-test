package main

import (
	"encoding/json"
	"fmt"
)

func testJson() {
	for i := 0; i < 1000; i++ {
		cat := Cat{Name: "Murz", Age: 8, Color: "Green"}
		catJson, err := json.Marshal(cat)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(catJson))
	}
}
