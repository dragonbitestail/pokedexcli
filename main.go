package main

import (
	"os"
	"strconv"
	"time"
	"github.com/dragonbitestail/pokedexcli/internal/cache"
	"github.com/dragonbitestail/pokedexcli/internal/logging"
)

var logr = ilogger.GetLogger()

var cmdMap = map[string]cliCommand {
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
		description: "List Pokemon in given area: explore <map_area_name>",
		callback: commandExplore,
	},
	"catch": {
		name: "catch",
		description: "Attempt to Catch named Pokemon: catch <pokemon_name|id>",
		callback: commandCatch,
	},
	"inspect": {
		name: "inspect",
		description: "List details for a Pokemon: inspect <pokemon_name>",
		callback: commandInspect,
	},

}

// Force help items to print in predetermined order:
var helpKeyOrder = []string{"help", "exit", "map", "mapb", "explore", "catch", "inspect"}
var cfg = config{
	reg: cmdMap,
	helpOrder: helpKeyOrder,
	rootAPI: "https://pokeapi.co/api/v2",
	pokemonEP: "/pokemon/",
	//locationsAPI: SEE main()
	mapNextURL: "null",
	mapBackURL: "null",
	//pokeCache: SEE main()
}

func main() {
	logr.Info("Starting Pokedex REPL::...")

	//locationsEndPoint := "https://pokeapi.co/api/v2/location-area/"
	locationsEndPoint := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	//locationsEndPoint := "https://pokeapi.co/api/v2/location-area/BAD"

	cacheDuration := time.Second * 200
	cacheSecs, ok := os.LookupEnv("POKEDEX_CACHE_DUR_SECS")
	if ok {
		i, err := strconv.Atoi(cacheSecs)
    if err != nil {
        panic(err)
    }
		cacheDuration = time.Second * time.Duration(i)
	}
	cachePokeAPI := pokecache.NewCache(time.Duration(cacheDuration))

	cfg.locationsAPI = locationsEndPoint
	cfg.pokeCache = cachePokeAPI

	startRepl(&cfg)
}
