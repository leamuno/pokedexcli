package main

import (
	"time"

	"github.com/leamuno/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Hour)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
