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
		return fmt.Sprintf("")
	case King:
		return fmt.Sprintf("King")
	case Queen:
		return fmt.Sprintf("Queen")
	case Rook:
		return fmt.Sprintf("Rook")
	case Bishop:
		return fmt.Sprintf("Bishop")
	case Knight:
		return fmt.Sprintf("Knight")
	case Pawn:
		return fmt.Sprintf("Pawn")
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
