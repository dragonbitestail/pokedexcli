package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var lowered []string

	for _, word := range strings.Fields(text) {
		lowered = append(lowered, strings.ToLower(word))
	}

	return lowered
}

func startRepl(cfg *config){

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		line := scanner.Text()
		words := cleanInput(line)
		command := words[0]
		if cmd, ok := cfg.reg[command]; ok {
			if err := cmd.callback(cfg); err != nil {
				logr.Error(fmt.Sprintf("command `%s`, error: %v", command, err))
			}
		} else {
			fmt.Println("Unknown command", )
		}

	}

}
