package main

import (
    "fmt"
		"os"
//		"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	reg map[string]cliCommand
	helpOrder []string
}


func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

//var cmdMap map[string]cliCommand
//var helpKeyOrder []string
//var cfg config

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	//for _, cmdKeyName := range helpKeyOrder {
	for _, cmdKeyName := range cfg.helpOrder {
		//if cmd, ok := cmdMap[cmdKeyName]; ok {
		if cmd, ok := cfg.reg[cmdKeyName]; ok {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		} else {
			fmt.Printf("Bad helpKeyOrder name %s. Fix name in helpKeyOrder to match a key in cmdMap", cmdKeyName)
		}

	}	
	return nil
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
