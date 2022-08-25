package docs

import (
	"fmt"
	"github.com/unidoc/unioffice/spreadsheet"
	"log"
	"testing"
)

func TestCreateExcel(t *testing.T) {
	exc := spreadsheet.New()
	defer exc.Close()
	sheetOne := exc.AddSheet()
	for r := 0; r < 5; r++ {
		row := sheetOne.AddRow()
		for c := 0; c < 4; c++ {
			cell := row.AddCell()
			cell.SetString(fmt.Sprintf("row %d cell %d", r, c))
		}
	}
	if err := exc.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}
	err := exc.SaveToFile("../files/excel1.ods")
	if err != nil {
		log.Fatalln(err)
	}
}
