package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var version = "0.1.0"

var lf *os.File

func LOG(text string) {
	if lf != nil {
		lf.WriteString(fmt.Sprintf("%s\n", text))
	}
}

func main() {
	if os.Getenv("HELP_DEBUG") != "" {
		f, err := tea.LogToFile("debug.log", "help")
		if err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		}
		lf = f
		defer f.Close() // nolint:errcheck
	}
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Couldn't open config file:", err)
		os.Exit(1)
	}
	if _, err := tea.NewProgram(initialModel(*config)).Run(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}

// pause 5min warten, dann werbung
// 25min vor show mit werbung ende 3mal laufen 5min nach einlass starten
