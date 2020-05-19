package game

type connectfour struct {
	board   [][]int
	players []int
	player  int
}

func newConnectFour(width, height int, players ...int) connectfour {
	b := make([][]int, height)
	for i := range b {
		b[i] = make([]int, width)
	}
	return connectfour{
		board:   b,
		players: players,
		player:  players[0],
	}
}
