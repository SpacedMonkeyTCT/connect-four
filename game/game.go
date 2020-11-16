package game

import "fmt"

type ConnectFour struct {
	columns int
	rows    int
	board   [][]int
	player  int
}

func NewConnectFour(columns, rows int) *ConnectFour {
	b := make([][]int, columns)
	for x := range b {
		b[x] = make([]int, rows)
		for y := range b[x] {
			b[x][y] = 0
		}
	}
	return &ConnectFour{
		columns: columns,
		rows:    rows,
		board:   b,
		player:  1,
	}
}

func (cf *ConnectFour) MakeMove(column int) int {
	for y := 0; y < cf.rows; y++ {
		if cf.board[column-1][y] == 0 {
			cf.board[column-1][y] = cf.player
			cf.player = cf.player ^ 3
			fmt.Println(cf.board[column-1])
			return y + 1
		}
	}
	return 0
}

func (cf *ConnectFour) Columns() int {
	return cf.columns
}

func (cf *ConnectFour) Rows() int {
	return cf.rows
}

func (cf *ConnectFour) CurrentPlayer() int {
	return cf.player
}
