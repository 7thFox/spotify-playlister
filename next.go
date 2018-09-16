package main

import (
	"log"

	"github.com/zmb3/spotify"
)

var (
	currentTrackIndex = -1
	sortedTracks      map[spotify.ID]bool
	savedTracks       *spotify.SavedTrackPage
)

func initalizeSortedTracks(client *spotify.Client) {
	sortedTracks = map[spotify.ID]bool{}
	for _, p := range config.Playlists {
		var err error
		offset := 0
		opts := spotify.Options{Offset: &offset}
		for tracks, err := client.GetPlaylistTracksOpt(p, &opts, ""); err == nil && offset < tracks.Total; tracks, err = client.GetPlaylistTracksOpt(p, &opts, "") {

			for _, t := range tracks.Tracks {
				sortedTracks[t.Track.ID] = true
			}
			offset += len(tracks.Tracks)
		}

		if err != nil {
			log.Fatal("initalizeSortedTracks ", err)
		}
	}
}

// Next returns the next unsorted track that's been saved. Returns nil if none after the current.
func Next(client *spotify.Client) *spotify.SavedTrack {
	var err error
	if sortedTracks == nil {
		initalizeSortedTracks(client)
	}

	currentTrackIndex++
	opts := spotify.Options{Offset: &currentTrackIndex}

	if currentTrackIndex < *opts.Offset || currentTrackIndex >= *opts.Offset+20 || savedTracks == nil {
		savedTracks, err = client.CurrentUsersTracksOpt(&opts)
	}

	for ; err == nil && currentTrackIndex < savedTracks.Total; savedTracks, err = client.CurrentUsersTracksOpt(&opts) {
		for ; currentTrackIndex < savedTracks.Offset+len(savedTracks.Tracks); currentTrackIndex++ {
			t := savedTracks.Tracks[currentTrackIndex-savedTracks.Offset]
			if !sortedTracks[t.ID] {
				return &t
			}
		}

		opts = spotify.Options{Offset: &currentTrackIndex}
	}
	if err != nil {
		log.Fatal("next ", err)
	}
	return nil
}

// Previous returns the previous unsorted track that's been saved. Returns nil if none before the current.
func Previous(client *spotify.Client) *spotify.SavedTrack {
	var err error
	if sortedTracks == nil {
		initalizeSortedTracks(client)
	}

	currentTrackIndex--
	offset := currentTrackIndex - 19
	if offset < 0 {
		offset = 0
	}

	opts := spotify.Options{Offset: &offset}

	if currentTrackIndex < *opts.Offset || currentTrackIndex >= *opts.Offset+20 || savedTracks == nil {
		savedTracks, err = client.CurrentUsersTracksOpt(&opts)
	}

	for ; err == nil && currentTrackIndex >= 0; savedTracks, err = client.CurrentUsersTracksOpt(&opts) {
		for ; currentTrackIndex >= savedTracks.Offset; currentTrackIndex-- {
			t := savedTracks.Tracks[currentTrackIndex-savedTracks.Offset]
			if !sortedTracks[t.ID] {
				return &t
			}
		}

		offset = currentTrackIndex - 19
		if offset < 0 {
			offset = 0
		}
		opts = spotify.Options{Offset: &offset}
	}
	if err != nil {
		log.Fatal("previous ", err)
	}
	return nil
}
