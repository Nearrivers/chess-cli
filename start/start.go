
package start

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mritd/bubbles/common"
	"github.com/mritd/bubbles/selector"
)

type RootModel struct {
	Sl selector.Model
}

type StartingPrompt struct {
	Name string
	Description string
	SelectionMessage string
}

func NewRootModel() RootModel {
	return RootModel{
		Sl: selector.Model{
			Data: []interface{}{
				StartingPrompt{Name: "Jouer", Description: "Démarrer une partie d'échecs", SelectionMessage: "C'est parti !"},
				StartingPrompt{Name: "Paramètres", Description: "Paramétrer le jeu d'échec", SelectionMessage: "Pas encore prêt..."},
				StartingPrompt{Name: "Quitter", Description: "Quitter l'application", SelectionMessage: "Au revoir!"},
			},
			PerPage: 3,
			HeaderFunc: selector.DefaultHeaderFuncWithAppend("Que souhaitez-vous faire ?"),
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
				return common.FontColor(s.(StartingPrompt).SelectionMessage, selector.ColorFinished) + "\n"
			},	
		},
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg {
	case common.DONE:
		return m, tea.Quit
	}

	_, cmd := m.Sl.Update(msg)
	return m, cmd
}

func (m RootModel) View() string {
	return m.Sl.View()
}

