package block

type position struct {
	row int32
	col int32
}

func newPosition(row, col int32) *position {
	return &position{
		row: row,
		col: col,
	}
}

func (p *position) GetRow() int32 {
	return p.row
}

func (p *position) GetCol() int32 {
	return p.col
}
