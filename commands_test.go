package main

import (
	"testing"
	"bytes"
	//"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"github.com/dragonbitestail/pokedexcli/internal/cache"
)


var cacheDuration = time.Second * time.Duration(200)
var cachePokeAPI = pokecache.NewCache(time.Duration(cacheDuration))


// WIP: Utility function for command func calls which output to stdout as a string
// Modified from: https://go.dev/play/p/PNqa5M8zo7
// a link from SO:
// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
func getStdout(fcall func(*config, string, []string) error, cfg *config, cmd string, args []string) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	fcall(cfg, cmd, args) // works fine here

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	return out
}



func TestHelpCommand(t *testing.T) {
	//t.Errorf("TODO: write this test!")
	//cfg := config{}
	out := getStdout(commandHelp, &cfg, "help", []string{} )

	// reading our temp stdout
	//fmt.Println("previous output size:", len(out))
	//fmt.Print(out)

	// asser the output of "help" command contains 3 standard lines of: welcome, usage, spacer-line
	// followed by len command lines.
	//got := strings.Count(out, "\n") + 1	// plus 1 to account for last line not ending in nl
	got := strings.Count(out, "\n")
	want := len(cmdMap) + 3
	if got != want {
		t.Errorf("help out lines %d, did not match expected lines %d. " +
			"For each command, there should be corresponding entry in helpKeyOrder",
			got, want)
	}

}

func TestMapAndMapbCommand(t *testing.T) {
	cfg.pokeCache = cachePokeAPI
	cfg.locationsAPI = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	_ = getStdout(commandMap, &cfg, "map", []string{})
	_ = getStdout(commandMap, &cfg, "map", []string{})

	got := getStdout(commandMapBack, &cfg, "map", []string{})

	want := `canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
`
	if got != want {
		t.Errorf("got\n%s but, wanted\n%s", got, want)
	}
}

func TestExploreCommand(t *testing.T) {
	cfg.pokeCache = cachePokeAPI

	//got := getStdout(commandExplore, &cfg, "explore", []string{"1"})
	got := getStdout(commandExplore, &cfg, "explore", []string{"canalave-city-area"})

	want := `Exploring canalave-city-area...
Found Pokemon:
	- tentacool
	- tentacruel
	- staryu
	- magikarp
	- gyarados
	- wingull
	- pelipper
	- shellos
	- gastrodon
	- finneon
	- lumineon
`

	if got != want {
		t.Errorf("got\n%s but, wanted\n%s", got, want)
	}

}
