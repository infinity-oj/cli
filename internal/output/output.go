package output

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func NewTable(headers ...interface{}) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New(headers...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.WithWriter(color.Output)

	return tbl
}
