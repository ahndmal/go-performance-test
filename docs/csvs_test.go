package docs

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestReadCsv(t *testing.T) {
	// /home/andrii/files/airbnb_last_review.csv
	bts, err := os.ReadFile("/home/andrii/files/airbnb_last_review.csv")
	if err != nil {
		log.Panicln("error when reading file")
	}
	fileContents := string(bts)
	lines := strings.Split(fileContents, "\n")
	strDate := strings.Split(lines[1], ",")[2]

	//May 21 2019
	parsedDate, err := time.Parse("MMM dd yyyy", strDate)
	if err != nil {
		return
	}
	fmt.Println(parsedDate)
}
