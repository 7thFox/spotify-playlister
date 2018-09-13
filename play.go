package main

import (
	"log"

	"github.com/zmb3/spotify"
)

// TogglePlay toggles between playing or pauseing the current track.
func TogglePlay(client *spotify.Client, track *spotify.SavedTrack) {
	curPlay, err := client.PlayerCurrentlyPlaying()
	if err != nil {
		log.Println(err)
	}

	if curPlay.Item.ID != track.ID {
		opts := spotify.PlayOptions{
			URIs: []spotify.URI{track.URI},
		}
		client.PlayOpt(&opts)
	} else {
		if curPlay.Playing {
			err = client.Pause()
		} else {
			err = client.Play()
		}
		if err != nil {
			log.Println(err)
		}
	}

}
