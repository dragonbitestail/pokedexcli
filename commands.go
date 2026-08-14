
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
	
  resp, err := getResponse(targetURL, reqGetC)
  if err != nil {
      return err
  }
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status: %s", resp.Status)

	}

	//decr, err := getResponseBodyDecoder(targetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	locMap := locationMap{}

	decr := json.NewDecoder(resp.Body)
	if err := decr.Decode(&locMap); err != nil {
		return err
	}

	cfg.mapNextURL = locMap.NextURL
	cfg.mapBackURL = locMap.PreviousURL

	printLocations(locMap.Results)
	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.mapBackURL == "null" {
		fmt.Println( "you're on the first page")
		return nil
	}

	targetURL := cfg.locationsAPI
	if cfg.mapBackURL != cfg.locationsAPI && strings.HasPrefix(cfg.mapBackURL, "http") {
		targetURL = cfg.mapNextURL
	}
	
  resp, err := getResponse(targetURL, reqGetC)
  if err != nil {
      return err
  }
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status: %s", resp.Status)

	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	locMap := locationMap{}

	decr := json.NewDecoder(resp.Body)
	if err := decr.Decode(&locMap); err != nil {
		return err
	}

	cfg.mapNextURL = locMap.NextURL
	cfg.mapBackURL = locMap.PreviousURL

	printLocations(locMap.Results)
	return nil

}


func printLocations(locs []location){
	//for _, item := range locMap.Results {
	for _, item := range locs {
		fmt.Printf("%s\n", item.Name)
	}

}


//func getResponseBodyDecoder(url string) (*json.Decoder, error) {
//  resp, err := getResponse(url, reqGetC)
//  if err != nil {
//      return nil, err
//  }
//	//defer resp.Body.Close()
//
//	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
//		return nil, fmt.Errorf("server returned status: %s", resp.Status)
//
//	}
//	
//	decr := json.NewDecoder(resp.Body)
//	return decr, nil
//}
