package cell

import (
	"distributary/cli/format"
	"fmt"
)

func New(value string) *Cell {
	return &Cell{Value: value}
}

func (column *Cell) Print() {
	fmt.Print(fmt.Sprintf("%s%s%s", column.Style.Color, column.Value, format.Colors.Reset))
	fmt.Print(multiplyRune(' ', column.Width-len(column.Value)))
}
