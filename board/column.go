package board

import (
	"github/nearrivers/chess-cli/piece"

	tea "github.com/charmbracelet/bubbletea"
)

type Column map[string]Square

type Square struct {
	Piece piece.Piece
}

func (s Square) Init() tea.Cmd {
	return nil
}

func (s Square) View() string {
	return ""
}
