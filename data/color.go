package data

type Color int

const (
	Empty Color = iota
	Black
	White
)

func (color Color) String() string {
	switch color {
	case Black:
		return "Black"
	case White:
		return "White"
	case Empty:
		return ""
	}
	return ""
}
