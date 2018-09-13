package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/zmb3/spotify"
)

var listOutput = ""

// List displays the index, name, and ID of the available sorting playlists as defined by config.json.
func List(client *spotify.Client) {
	if listOutput == "" { // cache the output, since it will remain the same
		for i, pid := range config.Playlists {
			p, err := client.GetPlaylist(pid)
			if err != nil {
				log.Fatal("list ", err)
			}
			listOutput = fmt.Sprintf("%s%d. %s\t(%s)\n", listOutput, i, p.Name, pid)
		}
	}
	fmt.Print(listOutput)
}

// Clear clears the terminal screen.
func Clear() {
	// https://pastebin.com/1M64T8PV

	var c *exec.Cmd
	var doClear = true

	switch runtime.GOOS {
	case "darwin":
	case "linux":
		c = exec.Command("clear")
	case "windows":
		c = exec.Command("cmd", "/c", "cls")
	default:
		// Msg(InfoMsg, "Clear function not supported on current OS\n")
		doClear = false
	}
	if doClear {
		c.Stdout = os.Stdout
		c.Run()
	}
}
