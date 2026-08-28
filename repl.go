package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func cleanInput(text string) []string {
	var lowered []string

	for _, word := range strings.Fields(text) {
		lowered = append(lowered, strings.ToLower(word))
	}

	return lowered
}

func startRepl(cfg *config, reader io.Reader, isTest bool) (error , exitVal) {

	scanner := bufio.NewScanner(reader)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		line := scanner.Text()
		logr.Debug("scanner.text", "line", line)
		words := cleanInput(line)
		if len(words) < 1 {
			continue
		}
		command := words[0]
		if cmd, ok := cfg.reg[command]; ok {
			err, exitState := cmd.callback(cfg, command, words[1:])
			if err != nil {
				logr.Error(fmt.Sprintf("command `%s`, error: %v", command, err))
				if isTest {
					return err, exitState
				}
			}
			if exitState.exit && ! isTest {
				return nil, exitState
			}

		} else {
			fmt.Println("Unknown command")
		}
		if isTest && exitState.exit {
			return nil, exitState
		}
	}

}
