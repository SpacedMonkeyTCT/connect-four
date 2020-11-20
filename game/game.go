package game

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
	}
	return &ConnectFour{
		columns: columns,
		rows:    rows,
		board:   b,
		player:  1,
	}
}

func (cf *ConnectFour) MakeMove(column int) int {
	for row := 0; row < cf.rows; row++ {
		if cf.board[column-1][row] == 0 {
			cf.board[column-1][row] = cf.player
			cf.swapPlayers()
			return row + 1
		}
	}
	return 0
}

func (cf *ConnectFour) swapPlayers() {
	cf.player = cf.player ^ 3
}

func (cf *ConnectFour) CheckForWin() int {
	for c := range cf.board {
		for r := range cf.board[c] {
			player := cf.board[c][r]
			if player == 0 {
				break
			}
			if cf.horizontalLineLen(c, r) >= 4 ||
				cf.verticalLineLen(c, r) >= 4 ||
				cf.diagonalUpLineLen(c, r) >= 4 ||
				cf.diagonalDownLineLen(c, r) >= 4 {
				return player
			}
		}
	}
	return 0
}

func (cf *ConnectFour) horizontalLineLen(col, row int) int {
	player := cf.board[col][row]
	l := 0
	for c := col; c < len(cf.board); c++ {
		if cf.board[c][row] != player {
			break
		}
		l++
	}
	return l
}

func (cf *ConnectFour) verticalLineLen(col, row int) int {
	player := cf.board[col][row]
	l := 0
	for r := row; r < len(cf.board[col]); r++ {
		if cf.board[col][r] != player {
			break
		}
		l++
	}
	return l
}

func (cf *ConnectFour) diagonalUpLineLen(col, row int) int {
	player := cf.board[col][row]
	l := 0
	for c, r := col, row; c < len(cf.board) && r < len(cf.board[c]); c, r = c+1, r+1 {
		if cf.board[c][r] != player {
			break
		}
		l++
	}
	return l
}

func (cf *ConnectFour) diagonalDownLineLen(col, row int) int {
	player := cf.board[col][row]
	l := 0
	for c, r := col, row; c < len(cf.board) && r >= 0; c, r = c+1, r-1 {
		if cf.board[c][r] != player {
			break
		}
		l++
	}
	return l
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
