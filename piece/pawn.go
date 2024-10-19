package piece

type Pawn struct {
	side Side
}

func (p Pawn) Display() string {
	if p.side == WHITE {
		return WhitePawn
	}

	return BlackPawn
}
