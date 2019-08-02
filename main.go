package main

import (
	"fmt"

	"github.com/gochess_learning/data"
)

func main() {
	chessBoard := data.CreateInitialChessBoard()

	game := data.Match{ChessBoard: chessBoard, Turn: data.White}
	fmt.Println(game)

}
