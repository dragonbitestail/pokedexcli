package main

//import (
//    "fmt"
//		"os"
//		"strings"
//)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	reg map[string]cliCommand
	helpOrder []string
}


func main() {
	cmdMap := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
	// Force help items to print in predetermined order:
	helpKeyOrder := []string{"help", "exit"}

	cfg := config{
		reg: cmdMap,
		helpOrder: helpKeyOrder,
	}

	startRepl(&cfg)
}
