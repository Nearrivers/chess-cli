package piece

type Side string

const (
	WHITE       Side = "white"
	Black       Side = "black"
	WhiteKink        = "\U00002654"
	WhiteQueen       = "\U00002655"
	WhiteRook        = "\U00002656"
	WhiteBishop      = "\U00002657"
	WhiteKnight      = "\U00002658"
	WhitePawn        = "\U00002659"
	BlackKing        = "\U0000265A"
	BlackQueen       = "\U0000265B"
	BlackRook        = "\U0000265C"
	BlackBishop      = "\U0000265D"
	BlackKnight      = "\U0000265E"
	BlackPawn        = "\U0000265F"
)

type Piece interface {
	Display() string
}
