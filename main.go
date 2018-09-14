package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/zmb3/spotify"
)

var (
	auth = spotify.NewAuthenticator(redirectURI,
		spotify.ScopeUserReadPrivate,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopeUserLibraryRead,
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadCurrentlyPlaying,
	)
	ch        = make(chan *spotify.Client)
	state     = "abc123"
	config    Config
	autoclear = true
)

func main() {
	InitConfig()
	client := GetClient()
	reader := bufio.NewReader(os.Stdin)
	Clear()
	currentSong := Next(client)
	for {
		fmt.Println("")
		List(client)
		fmt.Printf("\n\n\tCurrent Song: %s - %s\n", currentSong.Name, currentSong.Artists[0].Name)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("main ", err)
		}
		if autoclear {
			Clear()
		}
		currentSong = ParseCmd(cmd, client, currentSong)
	}

}
