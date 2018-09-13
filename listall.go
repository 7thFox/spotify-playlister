package main

import (
	"fmt"
	"log"

	"github.com/zmb3/spotify"
)

// Listall lists all the user's playlists in a format that's easily copy-able to the config
func Listall(client *spotify.Client) {
	plsts, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Fatal("listall ", err)
	}

	for _, p := range plsts.Playlists {
		fmt.Printf("\"%s\", // %s\n", p.ID, p.Name)
	}
}
