
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

	targetURL := cfg.locationsAPI
	if cfg.mapNextURL != cfg.locationsAPI && strings.HasPrefix(cfg.mapNextURL, "http") {
		targetURL = cfg.mapNextURL
	}
	
  resp, bodyBytes, err := getResponseWithBodyAsBytes(targetURL, reqGetC)
  if err != nil {
      return err
  }

	// Ensure we have a happy HTTP status:
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status: %s", resp.Status)
	}

	locMap := locationMap{
		PreviousURL: "null",
		NextURL: "null",
	}

	if err := json.Unmarshal(bodyBytes, &locMap); err != nil {
		return err
	}
	//logr.Debug("JSON unmarshled to locationMap =: ", "values", fmt.Sprintf("%+v", locMap))
	logr.Debug("JSON unmarshled to locationMap w/ select values =: ",
		"Count", locMap.Count,
		"PreviousURL", locMap.PreviousURL,
		"NextURL", locMap.NextURL)

	cfg.mapNextURL = locMap.NextURL
	cfg.mapBackURL = locMap.PreviousURL

	printLocations(locMap.Results)
	return nil
}

func commandMapBack(cfg *config) error {
	logr.Debug("cfg =: ", "values", fmt.Sprintf("%+v", cfg))

	if cfg.mapBackURL == "null" {
		fmt.Println( "you're on the first page")
		return nil
	}

	targetURL := cfg.locationsAPI
	if cfg.mapBackURL != cfg.locationsAPI && strings.HasPrefix(cfg.mapBackURL, "http") {
		targetURL = cfg.mapBackURL
	}

	resp, bodyBytes, err := getResponseWithBodyAsBytes(targetURL, reqGetC)
  if err != nil {
      return err
  }

	// Ensure we have a happy HTTP status:
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status: %s", resp.Status)
	}

	locMap := locationMap{
		PreviousURL: "null",
		NextURL: "null",
	}

	if err := json.Unmarshal(bodyBytes, &locMap); err != nil {
		return err
	}
	logr.Debug("JSON unmarshled to locationMap w/ select values =: ",
		"Count", locMap.Count,
		"PreviousURL", locMap.PreviousURL,
		"NextURL", locMap.NextURL)

	cfg.mapNextURL = locMap.NextURL
	cfg.mapBackURL = locMap.PreviousURL

	printLocations(locMap.Results)
	return nil

}

func printLocations(locs []location){
	for _, item := range locs {
		fmt.Printf("%s\n", item.Name)
	}

}
