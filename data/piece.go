package data

import "fmt"

type PieceType int

const (
	EmptyPiece PieceType = iota
	King
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func (pieceType PieceType) String() string {
	switch pieceType {
	case EmptyPiece:
		return ""
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	}
	return ""
}

type Piece struct {
	PieceType PieceType
	Color     Color
}

func (piece Piece) String() string {
	return fmt.Sprintf("%v %v", piece.Color, piece.PieceType)
}
