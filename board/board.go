package board

type Board struct {
	width  int
	height int
}

func New(width, height int) Board {
	return Board{
		width:  width,
		height: height,
	}
}

func (b Board) Width() int {
	return b.width
}

func (b Board) Height() int {
	return b.height
}
