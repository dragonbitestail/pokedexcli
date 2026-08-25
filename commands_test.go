package main

import (
	"testing"
	"bytes"
	//"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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

func TestExploreCommand(t *testing.T) {
	t.Errorf("TODO: write this test!")
}
