
package main

import (
		"encoding/json"
    "fmt"
		"os"
		"strings"
)

// ==========Map Location structs
type location struct {
	Name string `json:"name"`
	URL string  `json:"url"`
}
type locationMap struct {
	Count int `json:"count"`
	NextURL string  `json:"next"`
	PreviousURL string  `json:"previous"`
	Results []location  `json:"results"`

}


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

func commandMap(cfg *config) error {
	logr.Debug("commandMap() IN", "values", fmt.Sprintf("%+v", cfg))

	targetURL := cfg.locationsAPI
	if cfg.mapNextURL != cfg.locationsAPI && strings.HasPrefix(cfg.mapNextURL, "http") {
		targetURL = cfg.mapNextURL
	}
	logr.Info("commandMap()", "FinalTargetURL", targetURL)
	
	locMap, err := getLocationMap(cfg, targetURL)
	if err != nil {
		return err
	}

	printLocations(locMap.Results)
	logr.Debug("commandMap() OUT", "values", fmt.Sprintf("%+v", cfg))
	return nil
}

func commandMapBack(cfg *config) error {
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

	locMap, err := getLocationMap(cfg, targetURL)
	if err != nil {
		return err
	}

	printLocations(locMap.Results)
	logr.Debug("commandMapBack() OUT", "values", fmt.Sprintf("%+v", cfg))
	return nil
}

// Return unmarshled response as locationMap from either cached or HTTP
// retrieved pokeAPI
func getLocationMap(cfg *config, targetURL string) (*locationMap, error) {

	cacheBytes, ok := cfg.pokeCache.Get(targetURL)
	if ! ok {
		var err error
		cacheBytes, err = getResponseBytes(targetURL)
		if err != nil {
			return nil, err
		}
		logr.Info("getLocationMap(): adding HTTP response bytes to cache", "targetURL", targetURL)
		cfg.pokeCache.Add(targetURL, cacheBytes)
	}

	locMap := locationMap{
		PreviousURL: "null",
		NextURL: "null",
	}

	if err := json.Unmarshal(cacheBytes, &locMap); err != nil {
		return nil, err
	}
	logr.Info("getLocationMap(): JSON unmarshled to locationMap w/ select values",
		"Count", locMap.Count,
		"PreviousURL", locMap.PreviousURL,
		"NextURL", locMap.NextURL)

	if targetURL == cfg.locationsAPI {
		cfg.mapBackURL = "null"
	} else {
		cfg.mapBackURL = locMap.PreviousURL
	}
	cfg.mapNextURL = locMap.NextURL

	logr.Info("getLocationMap(): updated cfg back & next", "mapBackURL", cfg.mapBackURL, "mapNextURL", cfg.mapNextURL)
	return &locMap, nil 
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
