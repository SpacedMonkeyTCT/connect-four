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
	var y int
	valid := false
	fmt.Println(cf.board[column-1])
	for y = 0; y < cf.rows; y++ {
		if cf.board[column-1][y] == 0 {
			valid = true
			cf.board[column-1][y] = cf.player
			fmt.Println("Player", cf.player)
			cf.player = cf.player ^ 3
			fmt.Println("becomes", cf.player)
			break
		}
	}
	if !valid {
		return 0
	}
	return y + 1
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
