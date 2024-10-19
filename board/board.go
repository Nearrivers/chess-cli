package board

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/table"
)

var (
	OddSquareBackground = lipgloss.NewStyle().Background(lipgloss.Color("#fff"))
	EvenSquareBackground = lipgloss.NewStyle().Background(lipgloss.Color("#000"))
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
)

type BoardModel struct{
	table *table.Model
	textInput *textinput.Model
}

func (bm BoardModel) Init() tea.Cmd {
	// bm.table = table.New().
	// 	Border(lipgloss.NormalBorder()).
	// 	BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
	// 	StyleFunc(func(row, col int) lipgloss.Style {
	// 		var style lipgloss.Style
	// 		switch{
	// 			case row%2 == 0:
	// 			case col%2 == 0:
	// 				style = EvenSquareBackground
	// 			default:
	// 				style = OddSquareBackground
	// 		}
	// 		return style
	// 	})

	return nil
}

func (bm BoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	bm.table, cmd = bm.table.Update(msg)
	return bm, cmd
}

func (bm BoardModel) View() string {
	return fmt.Sprintf("%s\n\n%s", bm.table, bm.textInput.View())
}
