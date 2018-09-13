package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/zmb3/spotify"
)

var (
	auth   = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopeUserLibraryRead, spotify.ScopeUserModifyPlaybackState, spotify.ScopeUserReadCurrentlyPlaying)
	ch     = make(chan *spotify.Client)
	state  = "abc123"
	config Config
)

func main() {
	InitConfig()
	fmt.Println(config)
	client := GetClient()
	reader := bufio.NewReader(os.Stdin)
	currentSong := Next(client)

	for {
		Clear()
		List(client)
		fmt.Printf("\n\n\tCurrent Song: %s - %s\n", currentSong.Name, currentSong.Artists[0].Name)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("main ", err)
		}
		ParseCmd(cmd, client, currentSong)
	}

}
