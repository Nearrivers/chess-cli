package start

import (
	"fmt"
	"github/nearrivers/chess-cli/board"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mritd/bubbles/common"
	"github.com/mritd/bubbles/selector"
)

type sessionState int

const (
	selectorView sessionState = iota
	gameView
)

type PromptAction string

const (
	PLAY     PromptAction = "Jouer"
	SETTINGS PromptAction = "Paramètres"
	QUIT     PromptAction = "Quitter"
)

type RootModel struct {
	state sessionState
	Sl    selector.Model
	bm    board.BoardModel
}

type StartingPrompt struct {
	Name             PromptAction
	Description      string
	SelectionMessage string
}

func NewRootModel() *RootModel {
	rm := &RootModel{
		state: selectorView,
	}

	sl := selector.Model{
		Data: []interface{}{
			StartingPrompt{Name: PLAY, Description: "Démarrer une partie d'échecs", SelectionMessage: "C'est parti !"},
			StartingPrompt{Name: SETTINGS, Description: "Paramétrer le jeu d'échec", SelectionMessage: "Pas encore prêt..."},
			StartingPrompt{Name: QUIT, Description: "Quitter l'application", SelectionMessage: "Au revoir!"},
		},
		PerPage: 3,
		HeaderFunc: func(m selector.Model, obj interface{}, gdIndex int) string {
			return "Que souhaitez-vous faire ?"
		},
		SelectedFunc: func(m selector.Model, obj interface{}, gdIndex int) string {
			t := obj.(StartingPrompt)
			return common.FontColor(fmt.Sprintf("[%d] %s", gdIndex+1, t.Name), selector.ColorSelected)
		},
		UnSelectedFunc: func(m selector.Model, obj interface{}, gdIndex int) string {
			t := obj.(StartingPrompt)
			return common.FontColor(fmt.Sprintf(" %d. %s", gdIndex+1, t.Name), selector.ColorUnSelected)
		},
		FooterFunc: func(m selector.Model, obj interface{}, gdIndex int) string {
			t := m.Selected().(StartingPrompt)
			footerTpl := `
Description: %s`
			return common.FontColor(fmt.Sprintf(footerTpl, t.Description), selector.ColorFooter)
		},
		FinishedFunc: func(s interface{}) string {
			if s.(StartingPrompt).Name == PLAY {
				rm.state = gameView
			} else {

			}
			return common.FontColor(s.(StartingPrompt).SelectionMessage, selector.ColorFinished) + "\n\n\n"
		},
	}

	rm.Sl = sl
	return rm
}

func (m *RootModel) Init() tea.Cmd {
	return nil
}

func (m *RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch strings.ToLower(msg.String()) {
		case "q":
			return m, cmd
		}
	}

	switch m.state {
	case selectorView:
		newSlState, newCmd := m.Sl.Update(msg)
		m.Sl = *newSlState
		cmd = newCmd
	case gameView:
		if !m.bm.IsGameOn {
			newBoard := board.NewBoardModel()
			m.bm = newBoard
		}

		newBoardState, newCmd := m.bm.Update(msg)
		boardModel, ok := newBoardState.(board.BoardModel)
		if !ok {
			panic("could not perform assertion en board model")
		}
		m.bm = boardModel
		cmd = newCmd

	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *RootModel) View() string {
	switch m.state {
	case gameView:
		return m.bm.View()
	default:
		return m.Sl.View()
	}
}
