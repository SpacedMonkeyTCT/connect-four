package game

type ConnectFour struct {
	board   [][]int
	players []int
	player  int
}

func NewConnectFour(width, height int, players ...int) ConnectFour {
	b := make([][]int, height)
	for i := range b {
		b[i] = make([]int, width)
	}
	return ConnectFour{
		board:   b,
		players: players,
		player:  players[0],
	}
}
