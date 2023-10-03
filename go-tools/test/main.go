package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// var (
// 	colTitleIndex     = "Find"
// 	colTitleFirstName = "Replace"
// 	colTitleLastName  = "Number of replacements"
// 	rowHeader         = table.Row{colTitleIndex, colTitleFirstName, colTitleLastName}
// )

// func demoTableFeatures() {
// 	t := table.NewWriter()

//		//==========================================================================
//		// Append a few rows and render to console
//		//==========================================================================
//		// a row need not be just strings
//		t.AppendRow(table.Row{`Mikihiko, who knows Erika's approach to training "you want to learn techniqur - do it yourself," was slightly puzzled by these words of Leo.`, `Mikihiko, who knows Erika's approach to training was slightly puzzled by Leo's words. He said "you want to learn technique? Do it yourself."`, 2})
//		// all rows need not have the same number of columns
//		t.AppendRow(table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"})
//		// table.Row is just a shorthand for []interface{}
//		t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
//		// time to take a peek
//		t.SetAutoIndex(true)
//		t.AppendHeader(rowHeader)
//		t.SetCaption("Table with Auto-Indexing (columns-only).\n")
//		t.SetStyle(table.StyleLight)
//		t.SetCaption("Table using the style 'StyleLight'.\n")
//		fmt.Println(t.Render())
//		//┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
//		//│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
//		//├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//		//│   1 │ Arya       │ Stark     │   3000 │                             │
//		//│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
//		//│ 300 │ Tyrion     │ Lannister │   5000 │                             │
//		//├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//		//│     │            │ TOTAL     │  10000 │                             │
//		//└─────┴────────────┴───────────┴────────┴─────────────────────────────┘
//	}
func main() {
	// demoTableFeatures()
	// demoTableColors()
	data := [][]string{
		[]string{"Test", "Replace", "1"},
		[]string{"", "", ""},
		[]string{`Mikihiko, who knows Erika's approach to training "you want to learn techniqur - do it yourself," was slightly puzzled by these words of Leo.`, `Mikihiko, who knows Erika's approach to training was slightly puzzled by Leo's words. He said "you want to learn technique? Do it yourself."`, "2"},
		[]string{"----------------------------------", "", ""},
		[]string{"it is felt that this is not", "it felt like this was not", "3"},
		[]string{"", "", ""},
		[]string{"It seems that Tatsuya would become Tatsuya most definitely, for Miyuki she still needs some time.", "It seems that Tatsuya had been unphased by what was said earlier while Miyuki still needed time for things to settle in.", "4"},
		[]string{"", "", ""},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Find", "Replace", "Number of replacements"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
