Do a `go build` followed by a `docker build .`
Sample hidden service with tor that generates a new onion address on start up. if you already have a private key than you can copy the `hostname` and `privet_key` file it `/tor/service/` and change the entry point to only run the webserver binary
