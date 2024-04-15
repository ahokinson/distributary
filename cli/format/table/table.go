package table

import (
	"distributary/cli/format/table/cell"
	"fmt"
)

func New(initializers []Initializer) Table {
	var columns []*Column
	var headers []string
	for _, initializer := range initializers {
		columns = append(columns, &Column{Cells: []*cell.Cell{}, Style: initializer.Style})
		headers = append(headers, initializer.Header)
	}

	table := Table{
		Columns: columns,
		Rows:    []*Row{},
	}

	table.AddRow(headers)

	return table
}

func (table *Table) AddRow(values []string) {
	row := Row{Cells: []*cell.Cell{}}

	for i, value := range values {
		c := cell.New(value)
		table.Columns[i].AddCell(c)
		row.AddCell(c)
	}

	table.Rows = append(table.Rows, &row)
}

func (table *Table) Format() {
	padding := 2

	for _, column := range table.Columns {
		for _, c := range column.Cells {
			if len(c.Value) > column.Width {
				column.Width = len(c.Value)
			}
		}

		for _, c := range column.Cells {
			c.Width = column.Width + padding
		}
		column.ApplyStyle()
	}
}

func (table *Table) Print() {
	for _, r := range table.Rows {
		r.Print()
		fmt.Println()
	}
}
