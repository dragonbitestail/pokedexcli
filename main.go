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
	"pokedex": {
		name: "pokedex",
		description: "List all caught Pokemon: pokedex",
		callback: commandPokedex,
	},

}

// Force help items to print in predetermined order:
var helpKeyOrder = []string{"help", "exit", "map", "mapb", "explore", "catch", "inspect", "pokedex"}

var cacheDuration = time.Second * 200
var cachePokeAPI *pokecache.Cache

var cfg = config{
	reg: cmdMap,
	helpOrder: helpKeyOrder,
	rootAPI: "https://pokeapi.co/api/v2",
	pokemonEP: "/pokemon/",
	//locationsAPI: SEE main()
	mapNextURL: "null",
	mapBackURL: "null",
	//pokeCache: SEE init()
}


func init(){
	cacheSecs, ok := os.LookupEnv("POKEDEX_CACHE_DUR_SECS")
	if ok {
		i, err := strconv.Atoi(cacheSecs)
	  if err != nil {
	      panic(err)
	  }
	cacheDuration = time.Second * time.Duration(i)
}

	cachePokeAPI = pokecache.NewCache(time.Duration(cacheDuration))

	cfg.pokeCache = cachePokeAPI

}

func main() {
	logr.Info("Starting Pokedex REPL::...")

	//locationsEndPoint := "https://pokeapi.co/api/v2/location-area/"
	locationsEndPoint := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	cfg.locationsAPI = locationsEndPoint
	cfg.pokeCache = cachePokeAPI

	isTest := false
	startRepl(&cfg, os.Stdin, isTest)
}
