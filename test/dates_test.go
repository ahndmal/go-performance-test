package test

import (
	"fmt"
	"testing"
	"time"
)

func TestDates(t *testing.T) {
	//May 21 2019
	strDate := "May 21 2019"
	parsedDate, err := time.Parse("MMMM dd yyyy", strDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedDate)
}
