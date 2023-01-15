package docs

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"log"
	"testing"
)

func TestDocsCreation(t *testing.T) {
	doc := document.New()
	defer doc.Close()

	// First Table
	{
		table := doc.AddTable()
		// width of the page
		table.Properties().SetWidthPercent(100)
		// with thick borers
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

		row := table.AddRow()
		run := row.AddCell().AddParagraph().AddRun()
		run.AddText("Name")
		run.Properties().SetHighlight(wml.ST_HighlightColorBlack)
		row.AddCell().AddParagraph().AddRun().AddText("John Smith")
		row = table.AddRow()
		row.AddCell().AddParagraph().AddRun().AddText("Street Address")
		row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")
	}

	doc.AddParagraph() // break up the consecutive tables

	// Second Table
	{
		table := doc.AddTable()
		// 4 inches wide
		table.Properties().SetWidth(4 * measurement.Inch)
		borders := table.Properties().Borders()
		// thin borders
		borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

		row := table.AddRow()
		cell := row.AddCell()
		// column span / merged cells
		cell.Properties().SetColumnSpan(2)

		run := cell.AddParagraph().AddRun()
		run.AddText("Cells can span multiple columns")

		row = table.AddRow()
		cell = row.AddCell()
		cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
		cell.AddParagraph().AddRun().AddText("Vertical Merge")
		row.AddCell().AddParagraph().AddRun().AddText("")

		row = table.AddRow()
		cell = row.AddCell()
		cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
		cell.AddParagraph().AddRun().AddText("Vertical Merge 2")
		row.AddCell().AddParagraph().AddRun().AddText("")

		row = table.AddRow()
		row.AddCell().AddParagraph().AddRun().AddText("Street Address")
		row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")

	}
	doc.AddParagraph()
	err := doc.SaveToFile("../resources/doc_tables.odt")
	if err != nil {
		log.Fatalln(err)
	}
}
