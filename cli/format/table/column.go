package table

import (
	"distributary/cli/format/table/cell"
)

func (column *Column) AddCell(cell *cell.Cell) {
	column.Cells = append(column.Cells, cell)
}

func (column *Column) ApplyStyle() {
	for _, c := range column.Cells[1:] {
		c.Style = column.Style
	}
}
