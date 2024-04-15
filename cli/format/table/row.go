package table

import (
	"distributary/cli/format/table/cell"
	"fmt"
)

func (row *Row) AddCell(cell *cell.Cell) {
	row.Cells = append(row.Cells, cell)
}

func (row *Row) Print() {
	for _, c := range row.Cells {
		c.Print()
	}
	fmt.Print("\n")
}
