package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zmb3/spotify"
)

// stringListToIDList takes a space delineated string of indexes, and returns a list of playlist IDs.
func stringListToIDList(list string) []spotify.ID {
	strNums := strings.Split(list, " ")
	playlists := []spotify.ID{}
	for _, s := range strNums {
		if n, err := strconv.Atoi(s); err == nil && n >= 0 && n < len(config.Playlists) {
			playlists = append(playlists, config.Playlists[n])
		}
	}
	return playlists
}

var (
	lastTrack     spotify.ID
	lastPlaylists []spotify.ID
)

// Sort adds the current track to the playlists defined by a space delineated string of indexes.
func Sort(client *spotify.Client, track *spotify.SavedTrack, playlists string) {
	playlistsParsed := stringListToIDList(playlists)
	lastTrack = track.ID
	lastPlaylists = playlistsParsed

	for _, p := range playlistsParsed {
		if _, err := client.AddTracksToPlaylist(p, track.ID); err != nil {
			log.Printf("Failed to add \"%s\" to %s\n %v\n", track.Name, p, err)
		} else {
			fmt.Printf("Added \"%s\" to %s\n", track.Name, p)
		}
	}
}

// Unsort removes the last track added to a playlist from all the playlists it was added to in that command.
func Unsort(client *spotify.Client) {
	if lastTrack == "" || lastPlaylists == nil {
		fmt.Println("Nothing to undo")
		return
	}

	for _, p := range lastPlaylists {
		if _, err := client.RemoveTracksFromPlaylist(p, lastTrack); err != nil {
			log.Printf("Failed to remove \"%s\" to %s\n %v\n", lastTrack, p, err)
		} else {
			fmt.Printf("Removed \"%s\" to %s\n", lastTrack, p)
		}
	}
	lastTrack = ""
	lastPlaylists = nil
}
