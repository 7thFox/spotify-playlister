package main

import (
	"log"

	"github.com/zmb3/spotify"
)

var (
	currentTrackIndex = 0
	savedTracks       *spotify.SavedTrackPage
	sortedTracks      map[spotify.ID]bool
)

func initalizeSortedTracks(client *spotify.Client) {
	sortedTracks = map[spotify.ID]bool{}
	for _, p := range config.Playlists {
		tracks, err := client.GetPlaylistTracks(p)
		if err != nil {
			log.Fatal("initalizeSortedTracks ", err)
		}
		for _, t := range tracks.Tracks {
			sortedTracks[t.Track.ID] = true
		}
	}
}

// Next returns the next unsorted track that's been saved. Returns nil if none after the current.
func Next(client *spotify.Client) *spotify.SavedTrack {
	if savedTracks == nil {
		var err error
		savedTracks, err = client.CurrentUsersTracks()
		if err != nil {
			log.Fatal("next ", err)
		}
	}
	if sortedTracks == nil {
		initalizeSortedTracks(client)
	}

	currentTrackIndex++
	for ; currentTrackIndex < len(savedTracks.Tracks); currentTrackIndex++ {
		t := savedTracks.Tracks[currentTrackIndex]
		if !sortedTracks[t.ID] {
			return &t
		}
	}
	return nil
}

// Previous returns the previous unsorted track that's been saved. Returns nil if none before the current.
func Previous(client *spotify.Client) *spotify.SavedTrack {
	if savedTracks == nil {
		var err error
		savedTracks, err = client.CurrentUsersTracks()
		if err != nil {
			log.Fatal("previous ", err)
		}
	}
	if sortedTracks == nil {
		initalizeSortedTracks(client)
	}
	currentTrackIndex--
	for ; currentTrackIndex >= 0; currentTrackIndex-- {
		t := savedTracks.Tracks[currentTrackIndex]
		if !sortedTracks[t.ID] {
			return &t
		}
	}
	return nil
}
