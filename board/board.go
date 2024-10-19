package board

import (
	"fmt"
	"github/nearrivers/chess-cli/piece"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	OddSquareBackground  = lipgloss.NewStyle().Background(lipgloss.Color("#fff"))
	EvenSquareBackground = lipgloss.NewStyle().Background(lipgloss.Color("#000"))
	baseStyle            = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240"))
)

type Board struct {
	// Vu que l'on utilise les coordonnées du style Nc3 pour dire comment on bouge
	// les pièces, j'utilise un map de map
	squares Row
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) String() string {
	return "\U00002654 \U00002655 \U00002656 \U00002657 \U00002658 \U00002659 \U0000265A \U0000265B \U0000265C \U0000265D \U0000265E \U0000265F"
}

type BoardModel struct {
	board     *Board
	viewport  viewport.Model
	TextInput textinput.Model
	// Indique la couleur qui joue
	player   piece.Side
	IsGameOn bool
}

func NewBoardModel() BoardModel {
	log.Println("nouvelle input")
	ti := textinput.New()
	ti.Placeholder = "Rentrez un coup"
	ti.Focus()
	ti.Width = 20

	return BoardModel{
		TextInput: ti,
		board:     NewBoard(),
		IsGameOn:  true,
	}
}

func (bm BoardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (bm BoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "p":
			return bm, tea.Quit
		// case "enter":
		// 	newBmState, newCmd := bm.Update(msg)
		// 	boardModel, ok := newBmState.(BoardModel)
		// 	if !ok {
		// 		panic("could not assert board model")
		// 	}
		// 	cmd = newCmd
		// 	bm = boardModel
		default:
			bm.TextInput, cmd = bm.TextInput.Update(msg)
		}
	}

	cmds = append(cmds, cmd)
	return bm, tea.Batch(cmds...)
}

func (bm BoardModel) View() string {
	return fmt.Sprintf("%s\n%s\n\n", bm.board, bm.TextInput.View())
}
