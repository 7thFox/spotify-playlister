# spotify-playlister
A small go utility to help in getting all of your Spotify library into playlists.

If you're like me, your Spotify consists of a bunch of random songs with about 20 sorted into their proper playlist.
That means, when you're craving a certain thing, it's either hoping suffle picks that, picking a single artist, or listening to the 5 songs you did manage to get in a playlist.

spotify-playlister hopes to solve that. It will go down your list of saved songs and find the ones you haven't sorted out yet. 
Don't recongnize it by title? No problem, because you can play the song on the device you're currently listening on within the console.

# Installaion

If you don't have it, you'll need to [install Go](https://golang.org/doc/install). After that, just the following command:

```bash
go install github.com/7thFox/spotify-playlister
```

This will create the .exe in your `go/bin` directory (eg: C:\Users\YourName\go\bin on Windows)

Then create a config (more below) in the same directory that you move the .exe, and you're good to go.

# Configuration

To run, 3 things are required in the `config.json` file. An example (`config_example.json`) is provided for you to copy over.

```json
{	
    "clientID": "<your client ID>",
    "clientSecret": "<your client secret>"	,
    "playlists": [
		"<playlist 1 ID>",
		"<playlist 2 ID>",
		"<playlist 3 ID>",
		"<playlist 4 ID>", 
		"<playlist 5 ID>", 
		"<playlist 6 ID>", 
		"<playlist 7 ID>", 
        "<playlist 8 ID>"
    ]
}
```

You'll need to [get a Spotify client ID/Secret](https://developer.spotify.com/my-applications/.), since I sadly cannot safely give one out (plus I'm sure it's against ToS).

After that, just run it the first time (you'll need to delete the blank playlist entries), and run `listall` to see all your Spotify Playlists.
Copy over all the ones you want to sort into, and remove the comments (Go is finicky about the format of JSON).

Now restart the program, and everything should be working!

# Note

I've only ran/tested this in Windows, so bugs may appear in other OS's.
