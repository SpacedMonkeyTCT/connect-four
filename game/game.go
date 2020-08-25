package game

import "fmt"

type ConnectFour struct {
	width   int
	height  int
	board   [][]int
	players []int
	pIndex  int
}

func NewConnectFour(width, height int, players ...int) *ConnectFour {
	b := make([][]int, width)
	for x := range b {
		b[x] = make([]int, height)
		for y := range b[x] {
			b[x][y] = 0
		}
	}
	return &ConnectFour{
		width:   width,
		height:  height,
		board:   b,
		players: players,
		pIndex:  0,
	}
}

func (cf *ConnectFour) MakeMove(column int) int {
	var y int
	valid := false
	fmt.Println(cf.board[column-1])
	for y = 0; y < cf.height; y++ {
		if cf.board[column-1][y] == 0 {
			valid = true
			cf.board[column-1][y] = cf.players[cf.pIndex]
			fmt.Println("Player", cf.players[cf.pIndex])
			cf.pIndex = (cf.pIndex + 1) % len(cf.players)
			fmt.Println("becomes", cf.players[cf.pIndex])
			break
		}
	}
	if !valid {
		return 0
	}
	return y + 1
}

func (cf *ConnectFour) CurrentPlayer() int {
	return cf.players[cf.pIndex]
}
