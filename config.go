package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/zmb3/spotify"
)

// Config holds the data stored in confg.json.
type Config struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`

	// The list of playlist ID's used for sorting
	Playlists []spotify.ID `json:"playlists"`
}

// InitConfig loads the json data from config.json into the global "config" variable.
func InitConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("initConfig fileread ", err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("initConfig ", err)
	}
}
