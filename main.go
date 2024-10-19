package main

import (
	"fmt"
	"github/nearrivers/chess-cli/start"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("impossible de lire le fichier de logs: %v", err)
	}
	defer f.Close()

	m := start.NewRootModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Il y a eu une erreur: %v", err)
		os.Exit(1)
	}
}
