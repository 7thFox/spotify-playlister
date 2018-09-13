package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8080/callback"

// GetClient returns the spotify client using the API info provided in config.json
func GetClient() *spotify.Client {

	auth.SetAuthInfo(config.ClientID, config.ClientSecret)
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	browser.OpenURL(url)

	// wait for auth to complete
	client := <-ch

	return client
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Login Completed!<script>window.close();</script>")
	ch <- &client
}
