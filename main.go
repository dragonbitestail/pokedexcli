package main

import (
	"log/slog"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	reg map[string]cliCommand
	helpOrder []string
	locationsAPI string
	mapNextURL string
	mapBackURL string
}

var handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    Level: getLogLevelFromEnv(),
})
var logr = slog.New(handler)

func main() {
	logr.Info("Starting Pokedex REPL::...")
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
		"map": {
			name: "map",
			description: "Display chunk of Pokemon world locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display chunk of Pokemon world locations backwards",
			callback: commandMapBack,
		},

	}
	// Force help items to print in predetermined order:
	helpKeyOrder := []string{"help", "exit"}

	locationsEndPoint := "https://pokeapi.co/api/v2/location-area/"
	//locationsEndPoint := "https://pokeapi.co/api/v2/location-area/BAD"

	cfg := config{
		reg: cmdMap,
		helpOrder: helpKeyOrder,
		locationsAPI: locationsEndPoint,
		mapNextURL: "null",
		mapBackURL: "null",
	}

	startRepl(&cfg)
}

func getLogLevelFromEnv() slog.Level {
    levelStr := os.Getenv("LOG_LEVEL")

    switch strings.ToLower(levelStr) {
    case "debug":
        return slog.LevelDebug
    case "info":
        return slog.LevelInfo
    case "warn":
        return slog.LevelWarn
    case "error":
        return slog.LevelError
    default:
        return slog.LevelWarn
    }
}
