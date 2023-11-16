package main

import (
	"time"

	"github.com/SpaskeISO/golang_cli_pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient          pokeapi.Client
	nextLoactionAreaURL    *string
	prevoisLocationAreaURL *string
	caughtPokemon          map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
