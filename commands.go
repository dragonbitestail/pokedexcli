
package main

import (
    "fmt"
		"os"
)


func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmdKeyName := range cfg.helpOrder {
		if cmd, ok := cfg.reg[cmdKeyName]; ok {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		} else {
			fmt.Printf("Bad helpKeyOrder name %s. Fix name in helpKeyOrder to match a key in cmdMap", cmdKeyName)
		}

	}	
	return nil
}
