package data

import (
	"fmt"
)

const boardSize = 8

type ChessBoard [boardSize][boardSize]Piece

func (board ChessBoard) String() string {
	board_str := ""
	for i := boardSize - 1; i >= 0; i-- {
		for j := 0; j < boardSize; j++ {
			board_str += fmt.Sprintf("%v\t", board[i][j])
		}
		board_str += "\n"
	}
	return board_str
}

func CreateInitialChessBoard() ChessBoard {
	chessBoard := ChessBoard{{}}

	chessBoard[1] = GeneratePawnRow(White)
	chessBoard[boardSize-2] = GeneratePawnRow(Black)

	chessBoard[0] = GenerateMultiPieceRow(White)
	chessBoard[boardSize-1] = GenerateMultiPieceRow(Black)

	return chessBoard
}

func GeneratePawnRow(color Color) [boardSize]Piece {
	pawnRow := [boardSize]Piece{}
	for i := 0; i < boardSize; i++ {
		pawnRow[i] = Piece{Color: color, PieceType: Pawn}
	}
	return pawnRow
}

func GenerateMultiPieceRow(color Color) [boardSize]Piece {
	row := [boardSize]Piece{}

	row[0] = Piece{Color: color, PieceType: Rook}
	row[7] = Piece{Color: color, PieceType: Rook}

	row[1] = Piece{Color: color, PieceType: Knight}
	row[6] = Piece{Color: color, PieceType: Knight}

	row[2] = Piece{Color: color, PieceType: Bishop}
	row[5] = Piece{Color: color, PieceType: Bishop}

	row[3] = Piece{Color: color, PieceType: Queen}
	row[4] = Piece{Color: color, PieceType: King}

	return row
}
