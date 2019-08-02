package data

import (
	"fmt"
)

type Match struct {
	ChessBoard ChessBoard
	Turn       Color
}

func (match Match) String() string {
	return fmt.Sprintf("Turn: %v\nBoard:\n%v", match.Turn, match.ChessBoard)
}
