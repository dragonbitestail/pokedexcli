package main

import (
	"testing"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)


// WIP: Utility function for command func calls which output to stdout as a string
// Modified from: https://go.dev/play/p/PNqa5M8zo7
// a link from SO:
// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
func getStdout(fcall func(*config, string, []string) error, cfg *config, cmd string, args []string) (string, error) {
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

	var ferr error
	ferr = fcall(cfg, cmd, args) // works fine here

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	return out, ferr
}

func TestHelpCommand(t *testing.T) {
	out, err := getStdout(commandHelp, &cfg, "help", []string{} )
	if err != nil {
		t.Errorf("%v", err)
	}

	// asser the output of "help" command contains 3 standard lines of: welcome, usage, spacer-line
	// followed by len command lines.
	got := strings.Count(out, "\n")
	want := len(cmdMap) + 3
	if got != want || err != nil {
		t.Errorf("help out lines %d, did not match expected lines %d. " +
			"For each command, there should be corresponding entry in helpKeyOrder",
			got, want)
	}

}

func TestMapAndMapbCommand(t *testing.T) {
	cfg.locationsAPI = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	var err error
	_, err = getStdout(commandMap, &cfg, "map", []string{})
	_, err = getStdout(commandMap, &cfg, "map", []string{})

	got, err := getStdout(commandMapBack, &cfg, "map", []string{})
	if err != nil {
		t.Errorf("%v", err)
	}

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

	//test_args := []string{"1"}
	test_args := []string{"canalave-city-area"}
	got, err := getStdout(commandExplore, &cfg, "explore", test_args)
	if err != nil {
		t.Errorf("%v", err)
	}

	//want := `Exploring canalave-city-area...
	want_tmpl := `Exploring %s...
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
	want := fmt.Sprintf(want_tmpl, test_args[0])
	if got != want {
		t.Errorf("got\n%s but, wanted\n%s", got, want)
	}

}

func TestCommandCatchAndInspectAndListPokedex(t *testing.T) {

	argsPoke := []string{"tentacool"}
	isCaughtMsg := fmt.Sprintf("%s was caught!\n", argsPoke[0])

	// Attempt to catch our Pokemon a reasonably number of times
	for i := 0; i < 10; i++ {
		caughtMsg, err := getStdout(commandCatch, &cfg, "catch", argsPoke)
		if err != nil {
			t.Errorf("%v", err)
		}
		if strings.Contains(caughtMsg, isCaughtMsg) {
			// Done. We have our Pokemon for inspect
			break
		}

		// Sleep a bit so we don't spam the server:
		throttle := time.Duration(rand.Intn(300) + 401) // 400 - 700
		fmt.Printf("CATCH-FAIL: %s. Retrying in %v", caughtMsg, throttle)
		time.Sleep(throttle * time.Millisecond)

	}
	got, err := getStdout(commandInspect, &cfg, "inspect", []string{"tentacool"})
	if err != nil {
		t.Errorf("%v", err)
	}

	want := `Name: tentacool
Height: 9
Weight: 455
Types:
	-hp: 40
	-attack: 40
	-defense: 35
	-special-attack: 50
	-special-defense: 100
	-speed: 70
Types:
	- water
	- poison
`
	if got != want {
		t.Errorf("got\n%sbut, wanted\n%s", got, want)
	}

	got, err = getStdout(commandPokedex, &cfg, "pokedex", []string{""})
	if err != nil {
		t.Errorf("%v", err)
	}

	want = `Your Pokedex:
	- tentacool
`
	if got != want {
		t.Errorf("got\n%sbut, wanted\n%s", got, want)
	}

}
