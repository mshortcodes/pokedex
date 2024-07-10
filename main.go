package main

import "github.com/mshortcodes/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(cfg)
}
