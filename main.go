package main

import (
	"time"
	"github.com/dragonbitestail/pokedexcli/internal/cache"
	"github.com/dragonbitestail/pokedexcli/internal/logging"
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
	pokeCache *pokecache.Cache
}


var logr = ilogger.GetLogger()

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
		"explore": {
			name: "explore",
			description: "List Pokemon in given area",
			callback: commandExplore,
		},

	}
	// Force help items to print in predetermined order:
	helpKeyOrder := []string{"help", "exit", "map", "mapb", "explore"}

	//locationsEndPoint := "https://pokeapi.co/api/v2/location-area/"
	locationsEndPoint := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	//locationsEndPoint := "https://pokeapi.co/api/v2/location-area/BAD"

	cachePokeAPI := pokecache.NewCache(time.Duration(time.Second * 200))

	cfg := config{
		reg: cmdMap,
		helpOrder: helpKeyOrder,
		locationsAPI: locationsEndPoint,
		mapNextURL: "null",
		mapBackURL: "null",
		pokeCache: cachePokeAPI,
	}

	startRepl(&cfg)
}
