package table

import (
	"distributary/cli/format/table/cell"
)

type Table struct {
	Rows    []*Row
	Columns []*Column
}

type Column struct {
	Cells []*cell.Cell
	Width int
	Style cell.Style
}

type Row struct {
	Cells []*cell.Cell
}

type Initializer struct {
	Header string
	Style  cell.Style
}
