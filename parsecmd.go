package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zmb3/spotify"
)

// ParseCmd executes the command given from the string provided.
func ParseCmd(cmd string, client *spotify.Client, currentSong *spotify.SavedTrack) {
	switch strings.TrimSpace(strings.SplitN(cmd, "", 2)[0]) {
	case "help":
	case "h":
	default:
		Help()
	case "list":
	case "l":
	case "listall":
	case "la":
		Listall(client)
	case "next":
	case "n":
		currentSong = Next(client)
		if currentSong == nil {
			currentSong = Previous(client)
			if currentSong == nil {
				fmt.Println("All Songs Sorted!")
				os.Exit(0)
			}
		}
	case "previous":
	case "p":
		currentSong = Previous(client)
		if currentSong == nil {
			currentSong = Next(client)
			if currentSong == nil {
				fmt.Println("All Songs Sorted!")
				os.Exit(0)
			}
		}
	case "toggleplay":
	case "play":
	case "stop":
	case "t":
		TogglePlay(client, currentSong)
	case "sort":
	case "s":
		Sort(client, currentSong, strings.TrimSpace(strings.SplitN(cmd, "", 2)[1]))
		currentSong = Next(client)
	case "unsort":
	case "undo":
	case "u":
		Unsort(client)
		currentSong = Previous(client)
	case "quit":
	case "q":
		os.Exit(0)
	}
}

// Help displays a message relating to the available commands.
func Help() {
	helptext := `help			[h]		this help message
list			[l]	lists playlists with ID's for copying into the whitelist
next			[n]		shows next song not in a playlist
previous		[p]		goes back a track
toggleplay		[t]		toggles playing the current song
sort {listid}	[s]		puts current song into playlist by numbers
unsort			[u]		removes the last track from the playlists it was added to
quit			[q]		exits the program
`
	fmt.Println(helptext)
}
