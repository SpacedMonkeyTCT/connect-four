package game

import "fmt"

type ConnectFour struct {
	width  int
	height int
	board  [][]int
	player int
}

func NewConnectFour(width, height int) *ConnectFour {
	b := make([][]int, width)
	for x := range b {
		b[x] = make([]int, height)
		for y := range b[x] {
			b[x][y] = 0
		}
	}
	return &ConnectFour{
		width:  width,
		height: height,
		board:  b,
		player: 1,
	}
}

func (cf *ConnectFour) MakeMove(column int) int {
	var y int
	valid := false
	fmt.Println(cf.board[column-1])
	for y = 0; y < cf.height; y++ {
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

func (cf *ConnectFour) Width() int {
	return cf.width
}

func (cf *ConnectFour) Height() int {
	return cf.height
}

func (cf *ConnectFour) CurrentPlayer() int {
	return cf.player
}
