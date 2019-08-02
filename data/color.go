package data

import (
	"fmt"
)

type Color int

const (
	Empty Color = iota
	Black
	White
)

func (color Color) String() string {
	switch color {
	case Black:
		return fmt.Sprintf("Black")
	case White:
		return fmt.Sprintf("White")
	case Empty:
		return fmt.Sprintf("")
	}
	return ""
}
