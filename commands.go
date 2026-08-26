package main

import (
		"encoding/json"
    "fmt"
		"math/rand"
		"os"
		"strings"
)


func commandExit(cfg *config, cmd string, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, cmd string, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmdKeyName := range cfg.helpOrder {
		if cmd, ok := cfg.reg[cmdKeyName]; ok {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		} else {
			return fmt.Errorf("Bad helpKeyOrder name %s. Fix name in helpKeyOrder to match a key in cmdMap", cmdKeyName)
		}

	}
	return nil
}
func commandInspect(cfg *config, cmd string, args []string) error {
	logr.Debug("commandInspect() IN", "values", fmt.Sprintf("%+v", cfg), "args", args)
	if len(args) < 1 {
		return fmt.Errorf("inspect command requires Pokemon name parameter: catch <pokemon_name>")
	}

	pokeIdentifier := args[0]
	pObj, ok := pokeReg[pokeIdentifier]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n",
		pObj.Name, pObj.Height, pObj.Weight)

	fmt.Println("Types:")
	for _, s := range pObj.Stats {
		fmt.Printf("\t-%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pObj.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}

	return nil
}

func commandCatch(cfg *config, cmd string, args []string) error {
	logr.Debug("commandCatch() IN", "values", fmt.Sprintf("%+v", cfg), "args", args)
	if len(args) < 1 {
		return fmt.Errorf("catch command requires Pokemon name parameter: catch <pokemon>")
	}
	// Get Pokemon requested and unmarshal data to a new Pokemon struct type
	pokeIdentifier := args[0]
	targetURL := cfg.rootAPI + cfg.pokemonEP + pokeIdentifier
	pokeObj, err := getObjFromAPI[Pokemon](cfg, targetURL)
	if err != nil {
		return err
	}
	logr.Info("Retrieved Pokemon", "pokemon", pokeObj.Name, "base_experience", pokeObj.BaseExperience)

	// Extract value to base capture percentage on.
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeObj.Name)

	roleResults := rand.Intn(pokeObj.BaseExperience) / 100
	if roleResults == 0 { // The higher the experience, the more often we should get 1, so success should be only when we get 0
		logr.Info("Pokemon caught", "name", pokeObj.Name)
		fmt.Println(pokeObj.Name, "was caught!")
		pokeReg[pokeObj.Name] = pokeObj
	}else {
		logr.Info("Pokemon got away", "name", pokeObj.Name)
		fmt.Println(pokeObj.Name, "escaped!")
	}

	return nil
}

func commandExplore(cfg *config, cmd string, args []string) error {
	logr.Debug("commandExplore() IN", "values", fmt.Sprintf("%+v", cfg), "args", args)
	if len(args) < 1 {
		return fmt.Errorf("explore command requires area parameter: explore <area>")
	}
	targetURL := "https://pokeapi.co/api/v2/location-area/" + args[0]
	logr.Info("commandExplore()", "FinalTargetURL", targetURL)

	areaMap, err := getObjFromAPI[PokeMapArea](cfg, targetURL)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	if len(areaMap.PokemonEncounters) > 0 {
		logr.Info("commandExplore()", "areaMap.Name", areaMap.Name)
		fmt.Println("Found Pokemon:")
	}
	for _, enctr := range areaMap.PokemonEncounters {
		fmt.Printf("\t- %s\n", enctr.Pokemon.Name)
	}

	return nil
}

func getObjFromAPI[T PokeAPI](cfg *config, targetURL string) (*T, error) {
	var obj T

	cacheBytes, ok := cfg.pokeCache.Get(targetURL)
	if ! ok {
		var err error
		cacheBytes, err = getResponseBytes(targetURL)
		if err != nil {
			return &obj, err
		}
		cfg.pokeCache.Add(targetURL, cacheBytes)
	}

	if err := json.Unmarshal(cacheBytes, &obj); err != nil {
		return &obj, err
	}
	logr.Info("getObjFromAPI(): JSON unmarshled to obj", "type", fmt.Sprintf("%T",obj))

	return &obj, nil
}

func commandMap(cfg *config, cmd string, args []string) error {
	logr.Debug("commandMap() IN", "values", fmt.Sprintf("%+v", cfg))

	targetURL := cfg.locationsAPI
	if cfg.mapNextURL != cfg.locationsAPI && strings.HasPrefix(cfg.mapNextURL, "http") {
		targetURL = cfg.mapNextURL
	}
	logr.Info("commandMap()", "FinalTargetURL", targetURL)

	locMap, err := getObjFromAPI[locationMap](cfg, targetURL)
	if err != nil {
		return err
	}

	if targetURL == cfg.locationsAPI {
		cfg.mapBackURL = "null"
	} else {
		cfg.mapBackURL = locMap.PreviousURL
	}
	cfg.mapNextURL = locMap.NextURL

	printLocations(locMap.Results)
	logr.Debug("commandMap() OUT", "values", fmt.Sprintf("%+v", cfg))
	return nil
}

func commandMapBack(cfg *config, cmd string, args []string) error {
	logr.Debug("commandMapBack() IN", "values", fmt.Sprintf("%+v", cfg))

	if cfg.mapBackURL == "null" {
		fmt.Println( "you're on the first page")
		return nil
	}

	targetURL := cfg.locationsAPI
	if cfg.mapBackURL != cfg.locationsAPI && strings.HasPrefix(cfg.mapBackURL, "http") {
		targetURL = cfg.mapBackURL
	}
	logr.Info("commandMapBack()", "FinalTargetURL", targetURL)

	//locMap, err := getLocationMap(cfg, targetURL)
	locMap, err := getObjFromAPI[locationMap](cfg, targetURL)
	if err != nil {
		return err
	}

	if targetURL == cfg.locationsAPI {
		cfg.mapBackURL = "null"
	} else {
		cfg.mapBackURL = locMap.PreviousURL
	}
	cfg.mapNextURL = locMap.NextURL

	printLocations(locMap.Results)
	logr.Debug("commandMapBack() OUT", "values", fmt.Sprintf("%+v", cfg))
	return nil
}

func getResponseBytes(targetURL string) ([]byte, error) {
	resp, bodyBytes, err := getBodyBytesWithResponseHTTP(targetURL, reqGetC)
  if err != nil {
      return nil, err
  }

	// Ensure we have a happy HTTP status:
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("server returned status: %s", resp.Status)
	}

	return bodyBytes, nil
}
func printLocations(locs []location){
	for _, item := range locs {
		fmt.Printf("%s\n", item.Name)
	}

}
